package admin

import (
	"fiber/app/system/admin/api"
	"fiber/app/system/admin/service"
	"github.com/gogf/gf/net/ghttp"
)

func Init(s *ghttp.Server) {
	service.Auth.LoadToken()
	s.Group("api/v1/admin", func(group *ghttp.RouterGroup) {
		api.System.LoadRouter(group)
		api.Manger.LoadRouter(group)
		api.Role.LoadRouter(group)
		api.Grade.LoadRouter(group)
		api.Vip.LoadRouter(group)
		api.User.LoadRouter(group)
		api.Authority.LoadRouter(group)
		api.Article.LoadRouter(group)
		api.Audio.LoadRouter(group)
		api.Video.LoadRouter(group)
		api.Resource.LoadRouter(group)
		api.Question.LoadRouter(group)
		api.Topic.LoadRouter(group)
		api.Comment.LoadRouter(group)
		api.Answer.LoadRouter(group)
		api.Edu.LoadRouter(group)
		api.Group.LoadRouter(group)
		api.Category.LoadRouter(group)
		api.Tag.LoadRouter(group)
		api.Report.LoadRouter(group)
		api.Media.LoadRouter(group)
		api.Upload.LoadRouter(group)
		api.Order.LoadRouter(group)
		api.Recharge.LoadRouter(group)
		api.Cash.LoadRouter(group)
		api.Card.LoadRouter(group)
		api.Verify.LoadRouter(group)
	})
}
