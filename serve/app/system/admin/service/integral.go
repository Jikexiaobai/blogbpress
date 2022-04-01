package service

import (
	"fiber/app/dao"
	"fiber/app/system/admin/shared"
	"fiber/library/redis"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
)

var Integral = new(integralService)

type integralService struct {
}

func (s *integralService) SetUserContentIntegral(tx *gdb.TX, userId int64) error {
	val, err := System.FindValue(shared.IntegralSetting)
	if err != nil {
		return err
	}
	j := gjson.New(val)
	var redisCom redis.Com
	// 设置次数
	redisCom.Key = shared.ContentCreateIntegralCount
	redisCom.Filed = gconv.String(userId)
	redisCom.Data = 1
	count, err := redisCom.INCHashField()
	postIntegral := gconv.Int64(j.Get("contentIntegral"))
	postCount := gconv.Int64(j.Get("contentCount"))
	if err != nil {
		return err
	}
	if count <= postCount {
		// 获取当前积分
		userIntegral, err := dao.SysUser.Value(dao.SysUser.Columns.Integral, dao.SysUser.Columns.UserId, userId)
		err = User.EditAccountIntegral(tx, userId, gconv.Int64(userIntegral)+postIntegral)
		if err != nil {
			return err
		}

	}
	return nil
}

func (s *integralService) SetUserCommentIntegral(tx *gdb.TX, userId int64) error {
	val, err := System.FindValue("IntegralSetting")
	if err != nil {
		return err
	}
	j := gjson.New(val)
	var redisCom redis.Com
	// 设置次数
	redisCom.Key = shared.CommentCreateIntegralCount
	redisCom.Filed = gconv.String(userId)
	redisCom.Data = 1
	count, err := redisCom.INCHashField()
	postIntegral := gconv.Int64(j.Get("commentIntegral"))
	postCount := gconv.Int64(j.Get("commentCount"))
	if err != nil {
		return err
	}
	if count <= postCount {
		// 获取当前积分
		userIntegral, err := dao.SysUser.Value(dao.SysUser.Columns.Integral, dao.SysUser.Columns.UserId, userId)
		err = User.EditAccountIntegral(tx, userId, gconv.Int64(userIntegral)+postIntegral)
		if err != nil {
			return err
		}

	}
	return nil
}

func (s *integralService) SetUserAnswerIntegral(tx *gdb.TX, userId int64) error {
	val, err := System.FindValue("IntegralSetting")
	if err != nil {
		return err
	}
	j := gjson.New(val)
	var redisCom redis.Com
	// 设置次数
	redisCom.Key = shared.AnswerCreateIntegralCount
	redisCom.Filed = gconv.String(userId)
	redisCom.Data = 1
	count, err := redisCom.INCHashField()
	postIntegral := gconv.Int64(j.Get("answerIntegral"))
	postCount := gconv.Int64(j.Get("answerCount"))
	if err != nil {
		return err
	}
	if count <= postCount {
		// 获取当前积分
		userIntegral, err := dao.SysUser.Value(dao.SysUser.Columns.Integral, dao.SysUser.Columns.UserId, userId)
		err = User.EditAccountIntegral(tx, userId, gconv.Int64(userIntegral)+postIntegral)
		if err != nil {
			return err
		}

	}
	return nil
}

func (s *integralService) SetUserReportIntegral(tx *gdb.TX, userId int64) error {
	val, err := System.FindValue("IntegralSetting")
	if err != nil {
		return err
	}
	j := gjson.New(val)
	var redisCom redis.Com
	// 设置次数
	redisCom.Key = shared.ReportIntegralCount
	redisCom.Filed = gconv.String(userId)
	redisCom.Data = 1
	count, err := redisCom.INCHashField()
	postIntegral := gconv.Int64(j.Get("reportIntegral"))
	postCount := gconv.Int64(j.Get("reportCount"))
	if err != nil {
		return err
	}
	if count <= postCount {
		// 获取当前积分
		userIntegral, err := dao.SysUser.Value(dao.SysUser.Columns.Integral, dao.SysUser.Columns.UserId, userId)
		err = User.EditAccountIntegral(tx, userId, gconv.Int64(userIntegral)+postIntegral)
		if err != nil {
			return err
		}

	}
	return nil
}
