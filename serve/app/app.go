package app

import (
	"fiber/app/dao"
	"fiber/app/middleware"
	"fiber/app/system/admin"
	"fiber/app/system/archive"
	"fiber/app/system/index"
	"fiber/app/task"
	_ "fiber/app/valid"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

func Run() {
	s := g.Server()
	s.Use(middleware.Cors)
	s.SetIndexFolder(true)
	s.AddStaticPath("/public", "./public")

	// 设置进程全局时区
	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	//初始化设置文件上传大小
	FileSetting, err := dao.SysConfig.
		Value(dao.SysConfig.Columns.ConfigValue,
			dao.SysConfig.Columns.ConfigKey, "FileSetting")
	if err != nil {
		panic(err)
	}
	FileSettingJson := gjson.New(FileSetting)
	fileSize := FileSettingJson.GetInt64("fileSize") * 1024 * 1021
	s.SetClientMaxBodySize(fileSize)

	// 业务系统初始化
	admin.Init(s)
	index.Init(s)
	archive.Init(s)
	task.Init()
	s.Run()
}
