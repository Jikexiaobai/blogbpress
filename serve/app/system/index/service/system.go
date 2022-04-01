package service

import (
	"fiber/app/dao"
	"fiber/app/system/index/dto"
	"fiber/app/system/index/result"
	"fiber/app/system/index/shared"
	lock_utils "fiber/app/tools/lock"
	"fiber/app/tools/response"
	"fiber/library/redis"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var System = new(systemService)

type systemService struct{}

//Home 获取首页内容
func (s *systemService) Home() (interface{}, response.ResponseCode) {
	var redisCom redis.Com
	redisCom.Key = "home"
	info := make(map[string]interface{})

	forumData, err := redisCom.GetString()
	if err != nil {
		return nil, response.CACHE_READ_ERROR
	}
	if forumData != nil {
		err := gconv.Struct(forumData, &info)
		if err != nil {
			return nil, response.INVALID
		}
		return info, response.SUCCESS
	}

	// 获取配置
	HomeDesign, err := Config.FindValue(shared.HomeDesign)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	HomeDesignJson := gjson.New(HomeDesign)
	var homeDesignList []*result.HomeDesignList
	err = HomeDesignJson.Structs(&homeDesignList)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	list := make([]map[string]interface{}, 0)
	for _, i := range homeDesignList {
		res := make(map[string]interface{})
		switch i.Style {
		case shared.WidgetSwipe:
			res["list"] = i.List
		case shared.WidgetAudio:
			// 获取音频
			audioList, err := Audio.SelectByHomeList(i.AudioIds)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			res["list"] = audioList
		case shared.WidgetVideo:
			// 获取投稿视频
			videoList, err := Video.SelectByHomeList(i.VideoIds)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			res["list"] = videoList
		case shared.WidgetResource:
			// 获取资源
			resourceList, err := Resource.SelectByHomeList(i.ResourceIds)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			res["list"] = resourceList
		case shared.WidgetCommunity:
			list := make([]interface{}, 0)
			// 获取问题
			questionList, err := Question.SelectByHomeList(i.QuestionIds)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			list = append(list, questionList)
			//获取动态

			topicList, err := Topic.SelectByHomeList(i.TopicIds)

			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			list = append(list, topicList)
			res["list"] = list
		case shared.WidgetVip:
			//获取会员
			vipList, err := Vip.SelectByHomeList(i.VipIds)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			//g.Dump(vipList)
			res["list"] = vipList
		case shared.WidgetArticle:
			// 获取文章
			articleList, err := Article.SelectByHomeList(i.ArticleIds)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			res["list"] = articleList
		case shared.WidgetEdu:
			// 获取课程
			eduList, err := Edu.SelectByHomeList(i.EduIds)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			res["list"] = eduList
		case shared.WidgetSearch:
			res["backgroundImage"] = i.BackgroundImage
			//获取热门搜索词
			searchList, code := s.HotSearch()
			if code != response.SUCCESS {
				return nil, code
			}
			res["list"] = searchList
		case shared.WidgetCustom:
			res["list"] = i.List
		case shared.WidgetImage:
			res["content"] = i.Content
		}
		res["height"] = i.Height
		res["title"] = i.Title
		res["showTitle"] = i.ShowTitle
		res["style"] = i.Style
		list = append(list, res)
	}
	info["list"] = list

	//存入缓存
	redisCom.Time = "180"
	redisCom.Data = info
	err = redisCom.SetStringEX()

	if err != nil {
		return nil, response.CACHE_SAVE_ERROR
	}

	return info, response.SUCCESS
}

//HFiveHome 获取H5首页内容
func (s *systemService) HFiveHome() (interface{}, response.ResponseCode) {
	var redisCom redis.Com
	redisCom.Key = "h5home"
	info := make(map[string]interface{})

	forumData, err := redisCom.GetString()
	if err != nil {
		return nil, response.CACHE_READ_ERROR
	}
	if forumData != nil {
		err := gconv.Struct(forumData, &info)
		if err != nil {
			return nil, response.INVALID
		}
		return info, response.SUCCESS
	}

	// 获取配置
	H5Design, err := Config.FindValue(shared.H5Design)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	HFiveDesignJson := gjson.New(H5Design)
	var h5DesignList []*result.HFiveDesignList
	err = HFiveDesignJson.Structs(&h5DesignList)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	list := make([]map[string]interface{}, 0)
	for _, i := range h5DesignList {
		res := make(map[string]interface{})
		switch i.Style {
		case shared.HFiveWidgetSwiper:
			res["list"] = i.List
		case shared.HFiveWidgetContent:
			contentList := make(map[string]interface{})
			if i.AudioIds != "" {
				// 获取音频
				audioList, err := Audio.SelectByHomeList(i.AudioIds)
				if err != nil {
					return nil, response.DB_READ_ERROR
				}
				contentList["audioList"] = audioList
			}
			if i.ResourceIds != "" {
				// 获取音频
				resourceList, err := Resource.SelectByHomeList(i.ResourceIds)
				if err != nil {
					return nil, response.DB_READ_ERROR
				}
				contentList["resourceList"] = resourceList
			}
			if i.VideoIds != "" {
				// 获取音频
				videoList, err := Video.SelectByHomeList(i.VideoIds)
				if err != nil {
					return nil, response.DB_READ_ERROR
				}
				contentList["videoList"] = videoList
			}
			res["list"] = contentList
		case shared.HFiveWidgetCommunity:
			list := make([]interface{}, 0)
			//// 获取问题
			//questionList, err := Question.SelectByHomeList(i.QuestionIds)
			//if err != nil {
			//	return nil, response.DB_READ_ERROR
			//}
			//list = append(list, questionList)
			//获取动态
			topicList, err := Topic.SelectByHomeList(i.TopicIds)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			list = append(list, topicList)
			res["list"] = list
		case shared.HFiveWidgetArticle:
			// 获取文章
			articleList, err := Article.SelectByHomeList(i.ArticleIds)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			res["list"] = articleList
		case shared.HFiveWidgetEdu:
			// 获取课程
			eduList, err := Edu.SelectByHomeList(i.EduIds)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			res["list"] = eduList
		case shared.HFiveWidgetImage:
			res["image"] = i.Image
			res["link"] = i.Link
			res["isPlatform"] = i.IsPlatform

			//获取热门搜索词
			searchList, code := s.HotSearch()
			if code != response.SUCCESS {
				return nil, code
			}
			res["list"] = searchList
		case shared.HFiveWidgetCustom:
			res["list"] = i.List
		}
		res["height"] = i.Height
		res["title"] = i.Title
		res["showTitle"] = i.ShowTitle
		res["style"] = i.Style
		list = append(list, res)
	}
	info["list"] = list

	//存入缓存
	redisCom.Time = "180"
	redisCom.Data = info
	err = redisCom.SetStringEX()

	if err != nil {
		return nil, response.CACHE_SAVE_ERROR
	}

	return info, response.SUCCESS
}

// Info 获取系统配置内容
func (s *systemService) Info() (*result.SystemInfo, response.ResponseCode) {

	var redisCom redis.Com
	redisCom.Key = shared.SystemInfo
	systemObj, err := redisCom.GetString()

	if err != nil {
		return nil, response.CACHE_READ_ERROR
	}
	if systemObj != nil {

		var systemInfo *result.SystemInfo
		err := gconv.Struct(systemObj, &systemInfo)
		if err != nil {
			return nil, response.INVALID
		}
		return systemInfo, response.SUCCESS
	}

	list, err := dao.SysConfig.
		Where(dao.SysConfig.Columns.ConfigKey+" IN(?)",
			[]string{shared.BaseSetting,
				shared.FileSetting, shared.PaySetting}).All()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	var base result.Base
	var file result.File
	var pay result.Pay
	for _, i := range list {
		if i.ConfigKey == shared.BaseSetting {
			err = gconv.Struct(i.ConfigValue, &base)
			if err != nil {
				return nil, response.INVALID
			}
		}
		if i.ConfigKey == shared.FileSetting {
			err = gconv.Struct(i.ConfigValue, &file)
			if err != nil {
				return nil, response.INVALID
			}
		}
		if i.ConfigKey == shared.PaySetting {
			err = gconv.Struct(i.ConfigValue, &pay)
			if err != nil {
				return nil, response.INVALID
			}
		}
	}
	info := result.SystemInfo{File: &file, Base: &base, Pay: &pay}

	// 写入缓存
	redisCom.Time = "1200"
	redisCom.Data = info
	err = redisCom.SetStringEX()
	if err != nil {
		return nil, response.CACHE_SAVE_ERROR
	}
	return &info, response.SUCCESS
}

// VipAndGradeList 获取系统配置内容
func (s *systemService) VipAndGradeList() (*result.SystemVipAndGradeList, response.ResponseCode) {
	// 获取 vip 列表
	gradeList, err := Grade.SelectList()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	vipList, err := Vip.SelectList()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	info := result.SystemVipAndGradeList{
		Grade: gradeList,
		Vip:   vipList,
	}
	return &info, response.SUCCESS
}

// SelectPostsList 过滤筛选内容
func (s *systemService) SelectPostsList(req *dto.QueryParam) (int, interface{}, response.ResponseCode) {
	var res []map[string]interface{}
	var total int

	if req.IsSearch && req.Title != "" {
		// 加入锁限制
		_, err := lock_utils.SetCount(shared.SearchCount+gconv.String(req.Related),
			shared.SearchLock+gconv.String(req.Related), 60, 10)
		if err != nil {
			return 0, nil, response.CACHE_SAVE_ERROR
		}

		var redisCom redis.Com
		redisCom.Key = shared.Search

		if redisCom.CheckExists() {
			redisCom.Data = req.Title
			err = redisCom.ZIncr()
			if err != nil {
				return 0, nil, response.CACHE_SAVE_ERROR
			}
		} else {
			redisCom.Time = "900"
			err = redisCom.SetExpire()
			if err != nil {
				return 0, nil, response.CACHE_SAVE_ERROR
			}
			redisCom.Data = req.Title
			err = redisCom.ZIncr()
			if err != nil {
				return 0, nil, response.CACHE_SAVE_ERROR
			}
		}

	}

	switch req.Module {
	case "":
		// 资源
		resourceCount, resourceList, err := Resource.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + resourceCount
		for _, i := range resourceList {
			m := gconv.Map(i)
			res = append(res, m)
		}

		// 视频
		videoCount, videoList, err := Video.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + videoCount
		for _, i := range videoList {
			m := gconv.Map(i)
			res = append(res, m)
		}

		// 音频
		audioCount, audioList, err := Audio.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + audioCount
		for _, i := range audioList {
			m := gconv.Map(i)
			res = append(res, m)
		}

		// 课程
		eduCount, eduList, err := Edu.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + eduCount
		for _, i := range eduList {
			m := gconv.Map(i)
			res = append(res, m)
		}

		// 文章
		articleCount, articleList, err := Article.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + articleCount
		for _, i := range articleList {
			m := gconv.Map(i)
			res = append(res, m)
		}

		// 帖子
		topicCount, topicList, err := Topic.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + topicCount
		for _, i := range topicList {
			m := gconv.Map(i)
			res = append(res, m)
		}

		// 问题
		questionCount, questionList, err := Question.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + questionCount
		for _, i := range questionList {
			m := gconv.Map(i)
			res = append(res, m)
		}
		//搜索用的
		if req.IsSearch {
			// 圈子
			groupCount, groupList, err := Group.SelectFilterList(req)
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			total = total + groupCount
			for _, i := range groupList {
				m := gconv.Map(i)
				res = append(res, m)
			}

			// 用户
			userCount, userList, err := User.SelectFilterList(req)
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			total = total + userCount
			for _, i := range userList {
				m := gconv.Map(i)
				res = append(res, m)
			}
		}
	case shared.Resource:
		// 资源
		count, list, err := Resource.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + count
		for _, i := range list {
			m := gconv.Map(i)
			res = append(res, m)
		}
	case shared.Video:
		// 视频
		count, list, err := Video.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + count
		for _, i := range list {
			m := gconv.Map(i)
			res = append(res, m)
		}
	case shared.Audio:
		// 音频
		count, list, err := Audio.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + count
		for _, i := range list {
			m := gconv.Map(i)
			res = append(res, m)
		}
	case shared.Edu:
		// 音频
		count, list, err := Edu.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + count
		for _, i := range list {
			m := gconv.Map(i)
			res = append(res, m)
		}
	case shared.Article:
		// 音频
		count, list, err := Article.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + count
		for _, i := range list {
			m := gconv.Map(i)
			res = append(res, m)
		}
	case shared.Topic:
		// 音频
		count, list, err := Topic.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + count
		for _, i := range list {
			m := gconv.Map(i)
			res = append(res, m)
		}
	case shared.Question:
		// 音频
		count, list, err := Question.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + count
		for _, i := range list {
			m := gconv.Map(i)
			res = append(res, m)
		}
	case shared.Group:
		// 圈子
		count, list, err := Group.SelectFilterList(req)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		total = total + count
		for _, i := range list {
			m := gconv.Map(i)
			res = append(res, m)
		}
	case shared.User:
		if req.IsSearch {
			// 音频
			count, list, err := User.SelectFilterList(req)
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			total = total + count
			for _, i := range list {
				m := gconv.Map(i)
				res = append(res, m)
			}
		}
	}

	if req.Module == "" && !req.IsSearch {
		switch req.Mode {
		case shared.ModeHot:
			tmpRes := garray.NewSortedArray(func(v1, v2 interface{}) int {
				tmpV1, _ := v1.(map[string]interface{})
				tmpV2, _ := v2.(map[string]interface{})
				tmp1Hot := tmpV1["hots"]
				tmp2Hot := tmpV2["hots"]
				if tmp1Hot.(int64) < tmp2Hot.(int64) {
					return 1
				}
				if tmp1Hot.(int64) > tmp2Hot.(int64) {
					return -1
				}

				return 0
			})
			for _, i := range res {
				tmpRes.Add(i)
			}
			return total, tmpRes, response.SUCCESS
		case shared.ModeNew:
			tmpRes := garray.NewSortedArray(func(v1, v2 interface{}) int {
				tmpV1, _ := v1.(map[string]interface{})
				tmpV2, _ := v2.(map[string]interface{})
				tmp1CreateTime := gtime.New(tmpV1["createTime"])
				tmp2CreateTime := gtime.New(tmpV2["createTime"])
				if tmp1CreateTime.Timestamp() < tmp2CreateTime.Timestamp() {
					return 1
				}
				if tmp1CreateTime.Timestamp() > tmp2CreateTime.Timestamp() {
					return -1
				}

				return 0
			})
			for _, i := range res {
				tmpRes.Add(i)
			}
			return total, tmpRes, response.SUCCESS
		}
	}

	return total, res, response.SUCCESS
}

// HotSearch 获取热搜关键词
func (s *systemService) HotSearch() ([]string, response.ResponseCode) {
	var redisCom redis.Com
	redisCom.Key = shared.Search
	hotList := redisCom.ZRange()
	return hotList, response.SUCCESS
}

// HotUser 获取热门用户
func (s *systemService) HotUser() ([]*result.UserHotList, response.ResponseCode) {
	var redisCom redis.Com
	redisCom.Key = shared.UserHot
	hotList := redisCom.GetRandSet(10)
	var list []*result.UserHotList
	err := dao.SysUser.
		Fields(dao.SysUser.Columns.UserId,
			dao.SysUser.Columns.NickName,
			dao.SysUser.Columns.Avatar,
		).
		Where(dao.SysUser.Columns.UserId+" IN(?)", hotList).
		Structs(&list)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	return list, response.SUCCESS
}
