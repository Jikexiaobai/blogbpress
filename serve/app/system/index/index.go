package index

import (
	"fiber/app/system/index/api"
	"fiber/app/system/index/service"
	"github.com/gogf/gf/net/ghttp"
)

func Init(s *ghttp.Server) {
	service.Auth.LoadToken()
	//ws.SetUp(s)
	s.Group("api/v1/web", func(group *ghttp.RouterGroup) {
		api.Auth.LoadRouter(group)
		api.System.LoadRouter(group)
		api.Account.LoadRouter(group)
		api.Upload.LoadRouter(group)
		api.Cash.LoadRouter(group)
		api.Order.LoadRouter(group)
		api.Recharge.LoadRouter(group)
		api.Article.LoadRouter(group)
		api.Audio.LoadRouter(group)
		api.Edu.LoadRouter(group)
		api.Resource.LoadRouter(group)
		api.Video.LoadRouter(group)
		api.Group.LoadRouter(group)
		api.Topic.LoadRouter(group)
		api.Question.LoadRouter(group)
		api.Comment.LoadRouter(group)
		api.Answer.LoadRouter(group)
		api.User.LoadRouter(group)

		api.Notice.LoadRouter(group)
		//api.Pay.LoadRouter(group)
		//api.Report.LoadRouter(group)
	})
}
