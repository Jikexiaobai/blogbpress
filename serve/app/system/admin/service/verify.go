package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/system/admin/shared"
	"fiber/app/tools/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

var Verify = new(verifyService)

type verifyService struct{}

// SelectList 查询列表
func (s *verifyService) SelectList(req *dto.VerifyQuery) (int, []*result.VerifyList, response.ResponseCode) {
	model := dao.SysVerify.SysVerifyDao.Order(dao.SysVerify.Columns.UpdateTime + " desc")
	if req.Status != 0 {
		model = model.Where(dao.SysVerify.Columns.Status, req.Status)
	}
	if req.Name != "" {
		name := gstr.HideStr(req.Name, 50, "*")
		model = model.Where(dao.SysVerify.Columns.Name, name)
	}
	if req.Code != "" {
		code := gstr.HideStr(req.Code, 50, "*")
		model = model.Where(dao.SysVerify.Columns.Code, code)
	}
	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Page(req.Page, req.Limit)
	list, err := model.Fields(
		dao.SysVerify.Columns.VerifyId,
		dao.SysVerify.Columns.Name,
		dao.SysVerify.Columns.Code,
		dao.SysVerify.Columns.Mode,
		dao.SysVerify.Columns.Number,
		dao.SysVerify.Columns.CreateTime,
		dao.SysVerify.Columns.Status,
		dao.SysVerify.Columns.UserId,
	).All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.VerifyList
	for _, i := range list {
		var info *result.VerifyList
		err = gconv.Struct(i, &info)
		nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, i.UserId)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		info.NickName = gconv.String(nickName)
		res = append(res, info)
	}
	return total, res, response.SUCCESS
}

// Review 审核
func (s *verifyService) Review(req *dto.Review) (code response.ResponseCode) {
	info, err := dao.SysVerify.Fields(
		dao.SysVerify.Columns.UserId,
		dao.SysVerify.Columns.VerifyId,
		dao.SysVerify.Columns.Status).
		Where(dao.SysVerify.Columns.VerifyId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
	}
	tx, err := g.DB().Begin()
	if err != nil {
		return response.DB_TX_ERROR
	}
	defer func() {
		if code != response.SUCCESS {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var tmpIds []int64
	var tmpInfoList []*model.SysVerify
	for _, i := range info {
		if i.Status == 1 {
			tmpIds = append(tmpIds, i.VerifyId)
			tmpInfoList = append(tmpInfoList, i)
		}
	}

	if len(tmpIds) < 1 {
		return response.DB_READ_ERROR
	}
	_, err = tx.Update(dao.SysVerify.Table,
		g.Map{
			dao.SysVerify.Columns.Status: req.Status,
			dao.SysVerify.Columns.Remark: req.Remark,
		}, dao.SysVerify.Columns.VerifyId+" IN(?)", tmpIds)
	if err != nil {
		return response.DB_SAVE_ERROR
	}
	// 设置通知
	var noticeList []model.SysNotice
	for _, i := range tmpInfoList {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.Status = shared.NoticeStatusReview
		notice.Receiver = i.UserId
		notice.CreateTime = gtime.Now()
		notice.SystemType = shared.NoticeSysTemReview

		if req.Status == shared.StatusReviewed {
			notice.Content = "您的认证已通过审核"
		}
		if req.Status == shared.StatusRefuse {
			notice.Content = "您的认证未通过审核，原因：" + req.Remark
		}
		noticeList = append(noticeList, notice)
	}
	if len(noticeList) > 0 {
		_, err = tx.Insert(dao.SysNotice.Table, noticeList)
		if err != nil {
			return response.DB_TX_ERROR
		}
	}
	return response.SUCCESS
}

// Remove 删除
func (s *verifyService) Remove(idList []int64) response.ResponseCode {
	_, err := dao.SysVerify.Delete(dao.SysVerify.Columns.VerifyId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}
	return response.SUCCESS
}
