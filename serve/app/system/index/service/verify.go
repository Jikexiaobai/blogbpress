package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/index/dto"
	"fiber/app/system/index/result"
	"fiber/app/system/index/shared"
	"fiber/app/tools/response"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

var Verify = new(verifyService)

type verifyService struct{}

// SelectStatus 查询认证状态
func (s *verifyService) SelectStatus(userId int64) (*result.AccountVerifyStatusIsPayPrice, response.ResponseCode) {
	res, err := dao.SysVerify.Fields(dao.SysVerify.Columns.Status).
		Where(dao.SysVerify.Columns.UserId, userId).One()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	var verifyStatus result.AccountVerifyStatusIsPayPrice
	if res != nil {
		verifyStatus.Status = res.Status
	}
	verifyStatus.IsPay = s.CheckHasPay(userId)

	verifySetting, err := Config.FindValue(shared.UserSetting)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	j := gjson.New(verifySetting)
	price := gconv.Float64(j.Get("verifyPrice"))
	verifyStatus.Price = price
	return &verifyStatus, response.SUCCESS
}

// SelectInfo 查询认证信息
func (s *verifyService) SelectInfo(userId int64) (*result.VerifyInfo, response.ResponseCode) {
	info, err := dao.SysVerify.Where(dao.SysVerify.Columns.UserId, userId).One()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	if info == nil {
		return nil, response.NOT_FOUND
	}
	var rs *result.VerifyInfo
	err = gconv.Struct(info, &rs)
	if err != nil {
		return nil, response.INVALID
	}
	return rs, response.SUCCESS
}

// Create 创建认证信息
func (s *verifyService) Create(req *dto.VerifyCreate) (int, response.ResponseCode) {

	var entity model.SysVerify
	entity.Code = gstr.HideStr(req.Code, 50, "*")
	entity.Name = gstr.HideStr(req.Name, 50, "*")
	entity.Mode = req.Mode
	entity.Number = req.Number
	encryption, err := gmd5.Encrypt(req.Code + req.Name)
	if err != nil {
		return 0, response.INVALID
	}
	entity.Encryption = encryption
	entity.UserId = req.UserId

	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	entity.Status = 1

	_, err = dao.SysVerify.Save(entity)
	if err != nil {
		return 0, response.ADD_FAILED
	}
	return entity.Status, response.SUCCESS
}

// CheckIsVerify 检查是否认证
func (s *verifyService) CheckIsVerify(userId int64) bool {
	isVerify, err := dao.SysVerify.Value(dao.SysVerify.Columns.Status,
		dao.SysVerify.Columns.UserId, userId)
	if gconv.Int(isVerify) != 2 || err != nil {
		return false
	}

	return true
}

// CheckHasPay 检查是否设置认证服务费
func (s *verifyService) CheckHasPay(userId int64) bool {
	verifySetting, err := Config.FindValue(shared.UserSetting)
	if err != nil {
		return false
	}
	j := gjson.New(verifySetting)
	price := gconv.Float64(j.Get("verifyPrice"))
	if price == 0 {
		return true
	}
	return Order.CheckIsPay(userId, userId, "user", shared.OrderTypeNine)
}

// CheckHasInfo 检查是否创建了认证信息
func (s *verifyService) CheckHasInfo(code, name string) bool {
	encryption, _ := gmd5.Encrypt(code + name)
	count, err := dao.SysVerify.Where(dao.SysVerify.Columns.Encryption, encryption).Count()
	if err != nil || count > 0 {
		return true
	}
	return false
}
