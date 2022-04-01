package service

import (
	"fiber/app/dao"
	"fiber/app/system/index/shared"
	"fiber/library/redis"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var Integral = new(integralService)

type integralService struct {
}

// SetUserCommentIntegral 设置用户评论积分
func (s *integralService) SetUserCommentIntegral(redisCom redis.Com, tx *gdb.TX, userId int64) error {
	val, err := Config.FindValue(shared.IntegralSetting)
	if err != nil {
		return err
	}
	j := gjson.New(val)
	// 设置次数
	redisCom.Key = shared.CommentCreateIntegralCount
	redisCom.Filed = gconv.String(userId)
	redisCom.Data = 1
	count, err := redisCom.INCHashField()
	if err != nil {
		return err
	}
	postIntegral := gconv.Int64(j.Get("commentIntegral"))
	postCount := gconv.Int64(j.Get("commentCount"))
	if count <= postCount {
		// 获取当前积分
		user, err := dao.SysUser.
			Fields(dao.SysUser.Columns.Integral, dao.SysUser.Columns.Grade).
			Where(dao.SysUser.Columns.UserId, userId).
			One()
		if err != nil {
			return err
		}
		err = Account.EditIntegral(tx, userId, gconv.Int64(user.Integral)+postIntegral)
		if err != nil {
			return err
		}

		//	获取等级升级积分
		gradeList, err := dao.SysGrade.All()
		if err != nil {
			return err
		}
		tmpGradeId := user.Grade
		for _, i := range gradeList {
			if gconv.Int64(user.Integral)+postIntegral >= gconv.Int64(i.Integral) {
				user.Grade = gconv.Int64(i.GradeId)
			}
		}

		if user.Grade != tmpGradeId {
			err = Account.EditGrade(tx, userId, user.Grade)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

// SetUserAnswerIntegral 设置用户回答积分
func (s *integralService) SetUserAnswerIntegral(redisCom redis.Com, tx *gdb.TX, userId int64) error {
	val, err := Config.FindValue(shared.IntegralSetting)
	if err != nil {
		return err
	}
	j := gjson.New(val)
	// 设置次数
	redisCom.Key = shared.AnswerCreateIntegralCount
	redisCom.Filed = gconv.String(userId)
	redisCom.Data = 1
	count, err := redisCom.INCHashField()
	if err != nil {
		return err
	}
	postIntegral := gconv.Int64(j.Get("answerIntegral"))
	postCount := gconv.Int64(j.Get("answerCount"))
	if count <= postCount {
		// 获取当前积分
		user, err := dao.SysUser.
			Fields(dao.SysUser.Columns.Integral, dao.SysUser.Columns.Grade).
			Where(dao.SysUser.Columns.UserId, userId).
			One()
		if err != nil {
			return err
		}
		err = Account.EditIntegral(tx, userId, gconv.Int64(user.Integral)+postIntegral)
		if err != nil {
			return err
		}

		//	获取等级升级积分
		gradeList, err := dao.SysGrade.All()
		if err != nil {
			return err
		}
		tmpGradeId := user.Grade
		for _, i := range gradeList {
			if gconv.Int64(user.Integral)+postIntegral >= gconv.Int64(i.Integral) {
				user.Grade = gconv.Int64(i.GradeId)
			}
		}

		if user.Grade != tmpGradeId {
			err = Account.EditGrade(tx, userId, user.Grade)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

// SetUserLikeAndFavoriteIntegral 设置用户点赞和收藏积分
func (s *integralService) SetUserLikeAndFavoriteIntegral(redisCom redis.Com, tx *gdb.TX, userId int64) error {
	val, err := Config.FindValue(shared.Article)
	if err != nil {
		return err
	}
	j := gjson.New(val)
	// 设置次数
	redisCom.Key = shared.LikeFavoriteIntegralCount
	redisCom.Filed = gconv.String(userId)
	redisCom.Data = 1
	count, err := redisCom.INCHashField()
	if err != nil {
		return err
	}

	postIntegral := gconv.Int64(j.Get("likeFavoriteIntegral"))
	postCount := gconv.Int64(j.Get("likeFavoriteCount"))

	if count <= postCount {
		// 获取当前积分
		user, err := dao.SysUser.
			Fields(dao.SysUser.Columns.Integral, dao.SysUser.Columns.Grade).
			Where(dao.SysUser.Columns.UserId, userId).
			One()
		if err != nil {
			return err
		}
		_, err = tx.Update(dao.SysUser.Table, g.Map{
			dao.SysUser.Columns.Integral: gconv.Int64(user.Integral) + postIntegral,
		}, dao.SysUser.Columns.UserId, userId)
		if err != nil {
			return err
		}

		//	获取等级升级积分
		gradeList, err := dao.SysGrade.All()
		if err != nil {
			return err
		}
		tmpGradeId := user.Grade
		for _, i := range gradeList {
			if gconv.Int64(user.Integral)+postIntegral >= gconv.Int64(i.Integral) {
				user.Grade = gconv.Int64(i.GradeId)
			}
		}
		if user.Grade != tmpGradeId {
			err = Account.EditGrade(tx, userId, user.Grade)
			if err != nil {
				return err
			}

			_, err = tx.Update(dao.SysUser.Table, g.Map{
				dao.SysUser.Columns.Grade: user.Grade,
			}, dao.SysUser.Columns.UserId, userId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// SetUserFollowIntegral 设置用户关注积分
func (s *integralService) SetUserFollowIntegral(redisCom redis.Com, tx *gdb.TX, userId int64) error {
	val, err := Config.FindValue(shared.IntegralSetting)
	if err != nil {
		return err
	}
	j := gjson.New(val)
	// 设置次数
	redisCom.Key = shared.FollowIntegralCount
	redisCom.Filed = gconv.String(userId)
	redisCom.Data = 1
	count, err := redisCom.INCHashField()
	if err != nil {
		return err
	}

	postIntegral := gconv.Int64(j.Get("followIntegral"))
	postCount := gconv.Int64(j.Get("followCount"))

	if count <= postCount {
		// 获取当前积分
		user, err := dao.SysUser.
			Fields(dao.SysUser.Columns.Integral, dao.SysUser.Columns.Grade).
			Where(dao.SysUser.Columns.UserId, userId).
			One()
		if err != nil {
			return err
		}
		_, err = tx.Update(dao.SysUser.Table, g.Map{
			dao.SysUser.Columns.Integral: gconv.Int64(user.Integral) + postIntegral,
		}, dao.SysUser.Columns.UserId, userId)
		if err != nil {
			return err
		}

		//	获取等级升级积分
		gradeList, err := dao.SysGrade.All()
		if err != nil {
			return err
		}
		tmpGradeId := user.Grade
		for _, i := range gradeList {
			if gconv.Int64(user.Integral)+postIntegral >= gconv.Int64(i.Integral) {
				user.Grade = gconv.Int64(i.GradeId)
			}
		}
		if user.Grade != tmpGradeId {
			err = Account.EditGrade(tx, userId, user.Grade)
			if err != nil {
				return err
			}

			_, err = tx.Update(dao.SysUser.Table, g.Map{
				dao.SysUser.Columns.Grade: user.Grade,
			}, dao.SysUser.Columns.UserId, userId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
