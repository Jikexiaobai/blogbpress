package task

import (
	"fiber/app/dao"
	"fiber/app/system/index/shared"
	"fiber/library/redis"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
)

func loadSign() {
	saveSign()
}

func saveSign() {
	_, _ = gcron.Add("0 0 2 * * *", func() {
		var redisCom redis.Com
		redisCom.Key = shared.UserSignIn
		list := redisCom.GetHashAllField()
		if len(list) > 0 {

			for _, i := range list {
				redisCom.Filed = i
				info := redisCom.GetHashFieldString()
				integralRand := gstr.Split(info, "_")
				signInfo, err := dao.SysSign.Where(dao.SysSign.Columns.UserId, i).One()
				if err != nil {
					g.Dump(err.Error())
				}

				if signInfo != nil {
					if signInfo.CreateTime.AddDate(0, 0, 1).Format("Y-m-d") == gtime.New(integralRand[1]).Format("Y-m-d") {
						_, err := dao.SysSign.Where(dao.SysSign.Columns.UserId, i).Update(g.Map{
							dao.SysSign.Columns.Integral:   integralRand[0],
							dao.SysSign.Columns.Count:      signInfo.Count + 1,
							dao.SysSign.Columns.CreateTime: gtime.New(integralRand[1]),
						})
						if err != nil {
							g.Dump(err.Error())
						}
					} else {
						_, err = dao.SysSign.Save(g.Map{
							dao.SysSign.Columns.Integral:   integralRand[0],
							dao.SysSign.Columns.UserId:     i,
							dao.SysSign.Columns.Count:      1,
							dao.SysSign.Columns.CreateTime: gtime.New(integralRand[1]),
						})
						if err != nil {
							g.Dump(err.Error())
						}
					}
				} else {
					_, err = dao.SysSign.Save(g.Map{
						dao.SysSign.Columns.Integral:   integralRand[0],
						dao.SysSign.Columns.UserId:     i,
						dao.SysSign.Columns.Count:      1,
						dao.SysSign.Columns.CreateTime: gtime.New(integralRand[1]),
					})
					if err != nil {
						g.Dump(err.Error())
					}
				}
			}
			err := redisCom.Delete()
			if err != nil {
				g.Dump(err.Error())
			}
			redisCom.Key = shared.UserSignToday
			err = redisCom.Delete()
			if err != nil {
				g.Dump(err.Error())
			}

			redisCom.Key = shared.UserSignInList
			err = redisCom.Delete()
			if err != nil {
				g.Dump(err.Error())
			}
		}
	})
}
