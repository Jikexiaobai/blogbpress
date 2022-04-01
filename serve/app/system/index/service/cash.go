package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/index/dto"
	"fiber/app/system/index/result"
	"fiber/app/system/index/shared"
	lock_utils "fiber/app/tools/lock"
	"fiber/app/tools/response"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"github.com/shopspring/decimal"
)

var Cash = new(cashService)

type cashService struct {
}

// Create 创建提现记录
func (s *cashService) Create(req *dto.CashCreate) response.ResponseCode {

	// 加入锁限制
	_, err := lock_utils.SetCount(shared.CashCreateCount+gconv.String(req.UserId),
		shared.CashCreateLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	var entity model.SysCash
	err = gconv.Struct(req, &entity)
	if err != nil {
		return response.INVALID
	}

	entity.Mode = shared.StatusReview
	entity.Status = shared.StatusReview
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	Code := s.BuildNumber(entity.Mode, req.PayMethod, req.UserId)
	entity.Code = Code
	entity.UserId = req.UserId

	// 获取服务费
	PaySetting, err := Config.FindValue(shared.PaySetting)
	if err != nil {
		return response.DB_READ_ERROR
	}
	payJson := gjson.New(PaySetting)
	cashServicePercent := gconv.Float64(payJson.Get("cashServicePercent"))

	cashMoney := decimal.NewFromFloat(gconv.Float64(req.Money))

	serviceMoney := cashMoney.Mul(decimal.NewFromFloat(cashServicePercent))
	tmpMoney := cashMoney.Sub(serviceMoney)

	entity.CashMoney = gconv.Float64(tmpMoney)
	entity.Money = req.Money
	entity.ServiceMoney = gconv.Float64(serviceMoney)

	_, err = dao.SysCash.Save(entity)
	if err != nil {
		return response.ADD_FAILED
	}

	return response.SUCCESS
}

// SelectList 查询用户提现列表
func (s *cashService) SelectList(req *dto.CashQuery) (int, []*result.CashInfo, response.ResponseCode) {
	model := dao.SysCash.SysCashDao.Order(dao.SysCash.Columns.UpdateTime + " desc")
	if req.Status != 0 {
		model = model.Where(dao.SysCash.Columns.Status, req.Status)
	}

	if req.UserId != 0 {
		model = model.Where(dao.SysCash.Columns.UserId, req.UserId)
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
	var res []*result.CashInfo
	for _, i := range list {
		var contentInfo *result.CashInfo
		err = gconv.Struct(i, &contentInfo)
		res = append(res, contentInfo)
	}
	return total, res, response.SUCCESS
}

// CheckBalanceHasCash 检查余额是否充足
func (s *cashService) CheckBalanceHasCash(userId int64, Money float64) bool {
	balance, err := Account.SelectBalance(userId)
	if err != nil {
		return false
	}
	return balance > Money
}

// CheckCashMin 检查是否超过最低提现额度
func (s *cashService) CheckCashMin(Money float64) bool {
	PaySetting, err := Config.FindValue("PaySetting")
	if err != nil {
		return false
	}
	j := gjson.New(PaySetting)
	cashMin := gconv.Float64(j.Get("cashMin"))
	return Money > cashMin
}

// BuildNumber 业务类型 + 毫秒时间戳 + 6位随机数 + 用户id
func (s *cashService) BuildNumber(mode, paymentMethod int, userId int64) string {
	time := gconv.String(gtime.TimestampMilli())
	rand := grand.Digits(6)
	return "C" + gconv.String(mode) + time + rand + gconv.String(paymentMethod) + gconv.String(userId)
}
