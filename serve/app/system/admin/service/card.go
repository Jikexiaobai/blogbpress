package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/system/admin/shared"
	"fiber/app/tools/response"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
)

var Card = new(cradService)

type cradService struct {
}

// SelectList 查询列表
func (s *cradService) SelectList(req *dto.CardQuery) (int, []*result.CardList, response.ResponseCode) {
	model := dao.SysCard.SysCardDao.Order(dao.SysCard.Columns.CardId + " desc")
	if req.Status != 0 {
		model = model.Where(dao.SysCard.Columns.Status, req.Status)
	}

	if req.SecretKey != "" {
		model = model.Where(dao.SysCard.Columns.SecretKey, req.SecretKey)
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
	var res []*result.CardList
	for _, i := range list {
		var info *result.CardList
		err = gconv.Struct(i, &info)
		if i.UsedId != 0 {
			nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, i.UsedId)
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			info.NickName = gconv.String(nickName)
		}
		res = append(res, info)
	}
	return total, res, response.SUCCESS
}

// Create 创建
func (s *cradService) Create(req *dto.CardCreate) response.ResponseCode {

	var entity []model.SysCard
	time := gconv.String(gtime.TimestampMilli())
	rand := grand.Digits(6)

	createTime := gtime.Now()

	for i := 0; i < req.Count; i++ {
		tmpKey := "card" + time + rand + gconv.String(i) + gconv.String(req.Money)
		secretKey, err := gmd5.Encrypt(tmpKey)
		if err != nil {
			return response.FAILD
		}
		var tmpEntity model.SysCard
		tmpEntity.SecretKey = secretKey
		tmpEntity.CreateTime = createTime
		tmpEntity.Status = shared.StatusReview
		tmpEntity.Money = req.Money

		entity = append(entity, tmpEntity)
	}

	_, err := dao.SysCard.Insert(entity)
	if err != nil {
		return response.ADD_FAILED
	}
	return response.SUCCESS
}

// Remove 删除
func (s *cradService) Remove(idList []int64) response.ResponseCode {
	_, err := dao.SysCard.Delete(dao.SysCard.Columns.CardId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}
	return response.SUCCESS
}
