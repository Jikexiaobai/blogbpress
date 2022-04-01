package service

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/gsmtp"
	"github.com/gogf/gf/util/gconv"
)

var Email = new(emailService)

type emailService struct {
	Tos         string
	Subject     string
	Body        string
	ContentType string
}

func (e *emailService) Send() error {
	result, err := Config.FindValue("EmailOptions")
	if err != nil {
		return gerror.New("邮箱配置获取错误")
	}
	j := gjson.New(result)
	host := gconv.String(j.Get("host"))
	port := gconv.String(j.Get("port"))
	user := gconv.String(j.Get("user"))
	sender := gconv.String(j.Get("email"))
	pass := gconv.String(j.Get("pass"))
	address := host + ":" + port

	sm := gsmtp.New(address, user, pass)
	err = sm.SendMail(sender, e.Tos, e.Subject, e.Body, e.ContentType)

	if err != nil {
		return err
	}
	return nil
}
