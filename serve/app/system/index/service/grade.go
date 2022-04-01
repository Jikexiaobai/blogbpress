package service

import (
	"fiber/app/dao"
	"fiber/app/system/index/result"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

var Grade = new(gradeService)

type gradeService struct {
}

// SelectList 等级列表
func (s *gradeService) SelectList() ([]*result.Grade, error) {
	list, err := dao.SysGrade.All()
	if err != nil {
		return nil, err
	}
	var res []*result.Grade
	for _, i := range list {
		var info *result.Grade
		err = gconv.Struct(i, &info)
		if err != nil {
			return nil, err
		}
		postsModule := gjson.New(i.PostsModule)
		postsModuleList := gconv.Strings(postsModule.Value())
		if len(postsModuleList) > 0 {
			info.PostsModule = postsModuleList
		}

		commonAuth := gjson.New(i.CommonAuth)
		commonAuthList := gconv.Strings(commonAuth.Value())
		if len(commonAuthList) > 0 {
			info.CommonAuth = commonAuthList
		}
		res = append(res, info)
	}
	return res, nil
}

// SelectUserGrade 获取用户的等级
func (s *gradeService) SelectUserGrade(gradeId int64) (*result.UserGrade, error) {
	info, err := dao.SysGrade.Fields(
		dao.SysGrade.Columns.Title,
		dao.SysGrade.Columns.Icon,
	).Where(dao.SysGrade.Columns.GradeId, gradeId).One()
	if err != nil {
		return nil, err
	}
	var grade result.UserGrade
	err = gconv.Struct(info, &grade)
	if err != nil {
		return nil, err
	}
	return &grade, nil
}

// CheckHasCreateGroup 检查用户等级是否还能创建圈子
func (s *gradeService) CheckHasCreateGroup(userId int64) bool {
	//vipAndGrade,err := dao.
	userInfo, err := dao.SysUser.
		Fields(dao.SysUser.Columns.Grade).
		Where(dao.SysUser.Columns.UserId, userId).
		One()
	if err != nil {
		return false
	}

	count, err := dao.SysGroup.Where(dao.SysGroup.Columns.UserId, userId).Count()
	if err != nil {
		return false
	}

	var tmpCount int
	gradeCount, err := dao.SysGrade.
		Value(dao.SysGrade.Columns.CreateGroup, dao.SysGrade.Columns.GradeId, userInfo.Grade)

	if err != nil {
		return false
	}
	tmpCount = gconv.Int(gradeCount)

	return count <= tmpCount
}

// CheckHasCommon 检查用户通用权限
func (s *gradeService) CheckHasCommon(userId int64, auth string) bool {
	//vipAndGrade,err := dao.
	userInfo, err := dao.SysUser.
		Fields(dao.SysUser.Columns.Grade).
		Where(dao.SysUser.Columns.UserId, userId).
		One()
	if err != nil {
		return false
	}
	commonAuth, err := dao.SysGrade.
		Value(dao.SysGrade.Columns.CommonAuth,
			dao.SysGrade.Columns.GradeId, userInfo.Grade)
	if err != nil {
		return false
	}

	return gstr.Contains(gconv.String(commonAuth), auth)
}

// CheckHasPosts 检查用户投稿权限
func (s *gradeService) CheckHasPosts(userId int64, auth string) bool {
	//vipAndGrade,err := dao.
	userInfo, err := dao.SysUser.
		Fields(dao.SysUser.Columns.Grade).
		Where(dao.SysUser.Columns.UserId, userId).
		One()
	if err != nil {
		return false
	}
	postsAuth, err := dao.SysGrade.
		Value(dao.SysGrade.Columns.PostsModule,
			dao.SysGrade.Columns.GradeId, userInfo.Grade)
	if err != nil {
		return false
	}
	g.Dump(postsAuth)
	return gstr.Contains(gconv.String(postsAuth), auth)
}
