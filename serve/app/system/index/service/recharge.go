package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/index/dto"
	"fiber/app/system/index/result"
	"fiber/app/system/index/shared"
	lock_utils "fiber/app/tools/lock"
	"fiber/app/tools/response"
	pay_lib "fiber/library/pay"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
)

var Recharge = new(rechargeService)

type rechargeService struct {
}

// Create 创建订单
func (s *rechargeService) Create(req *dto.RechargeCreate) (string, response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.RechargeCreateCount+gconv.String(req.UserId),
		shared.RechargeCreateLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return "", response.CACHE_SAVE_ERROR
	}

	var entity model.SysRecharge
	err = gconv.Struct(req, &entity)
	if err != nil {
		return "", response.INVALID
	}
	// 生成流水号
	code := s.buildNumber(req.Mode, req.UserId)
	entity.Code = code
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	entity.UserId = req.UserId

	if req.Mode == 1 || req.Mode == 2 {
		entity.Status = shared.RechargeStatusThree
	}

	if req.Mode == 3 || req.Mode == 4 {
		entity.Status = shared.RechargeStatusOne
	}

	_, err = dao.SysRecharge.Save(entity)
	if err != nil {
		return "", response.ADD_FAILED
	}

	return code, response.SUCCESS
}

// Pay 支付订单
func (s *rechargeService) Pay(userId int64, code string) (*result.RechargerPayInfo, response.ResponseCode) {
	var qrInfo result.RechargerPayInfo
	entity, err := dao.SysRecharge.
		Where(dao.SysRecharge.Columns.UserId, userId).
		Where(dao.SysRecharge.Columns.Code, code).One()
	if err != nil || entity == nil {
		return nil, response.NOT_FOUND
	}

	payOptions, err := Config.FindValue(shared.PaySetting)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	// 获取支付存储设置
	j := gjson.New(payOptions)
	alyPay := gconv.Int(j.Get("alyPay"))
	weChatPay := gconv.Int(j.Get("weChatPay"))

	baseSetting, err := Config.FindValue(shared.BaseSetting)
	if err != nil {
		return nil, response.PAY_ERROR
	}
	// 获取文件存储设置
	base := gjson.New(baseSetting)
	url := gconv.String(base.Get("url"))
	switch entity.Mode {
	case shared.RechargeModeOne:
		switch alyPay {
		case 1:
			return nil, response.PAY_ERROR
		case 2:

			alyOptions, err := Config.FindValue(shared.AlyPayOptions)
			if err != nil {
				return nil, response.PAY_ERROR
			}

			// 获取文件存储设置
			j := gjson.New(alyOptions)
			appId := gconv.String(j.Get("appId"))
			privateKey := gconv.String(j.Get("privateKey"))

			alyCertPublicKey := Media.SelectMediaPath(gconv.String(j.Get("alyCertPublicKey")))
			alyRootCert := Media.SelectMediaPath(gconv.String(j.Get("alyRootCert")))
			appPublicKey := Media.SelectMediaPath(gconv.String(j.Get("appPublicKey")))
			payEngine := &pay_lib.AliPayBody{
				AppId:            appId,
				PrivateKey:       privateKey,
				NotifyUrl:        url + "/recharge/alipay/notice",
				Subject:          "用户充值",
				OutTradeNo:       entity.Code,
				TotalAmount:      entity.Money,
				AlyCertPublicKey: "." + alyCertPublicKey,
				AlyRootCert:      "." + alyRootCert,
				AppCertPublicKey: "." + appPublicKey,
			}
			pageUrl, err := payEngine.WebPay()
			if err != nil {
				return nil, response.PAY_ERROR
			}
			qrInfo = result.RechargerPayInfo{
				Mode:   entity.Mode,
				QrCode: pageUrl,
				Code:   entity.Code,
			}
			return &qrInfo, response.SUCCESS
		}
	case shared.RechargeModeTwo:
		switch weChatPay {
		case 1:
			return nil, response.PAY_ERROR
		case 2:
			return nil, response.PAY_ERROR
		}
	}
	return nil, response.PAY_ERROR
}

// SelectList 查询用户订单列表
func (s *rechargeService) SelectList(req *dto.RechargeQuery) (int, []*result.RechargerInfo, response.ResponseCode) {
	model := dao.SysRecharge.SysRechargeDao.Order(dao.SysRecharge.Columns.CreateTime, "desc")
	if req.Status != 0 {
		model = model.Where(dao.SysRecharge.Columns.Status, req.Status)
	}

	if req.UserId != 0 {
		model = model.Where(dao.SysRecharge.Columns.UserId, req.UserId)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Page(req.Page, req.Limit)
	list, err := model.All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.RechargerInfo
	for _, i := range list {
		var contentInfo result.RechargerInfo
		contentInfo.Code = i.Code
		contentInfo.Mode = i.Mode
		contentInfo.Money = i.Money
		contentInfo.Status = i.Status
		contentInfo.Remark = i.Remark
		contentInfo.CreateTime = i.CreateTime
		if i.Mode == shared.RechargeModeThree {
			money, err := dao.SysCard.Value(dao.SysCard.Columns.Money, dao.SysCard.Columns.SecretKey, i.CardKey)
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			contentInfo.Money = gconv.Float64(money)
		}
		res = append(res, &contentInfo)
	}

	return total, res, response.SUCCESS
}

//BuildNumber 业务类型 + 毫秒时间戳 + 6位随机数 + 用户id
func (s *rechargeService) buildNumber(mode int, userId int64) string {
	time := gconv.String(gtime.TimestampMilli())
	rand := grand.Digits(6)
	return "C" + time + rand + gconv.String(mode) + gconv.String(userId)
}
