package utils

import (
	"github.com/gogf/gf/encoding/gcharset"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/grand"
	"path/filepath"
	"regexp"
)

// Ext 获取文件后缀名
func Ext(path string) string {
	preg := `^.[a-zA-Z0-9]*`
	ext := filepath.Ext(path)
	reg, err := regexp.Compile(preg)
	if err != nil {
		return ""
	}

	return reg.FindString(ext)
}

// 返回一个随机字符串
func RandS(i int) string {
	randStr := grand.S(i)
	return randStr
}

// GetCityByIp 获取ip所在地址
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}

	if ip == "[::1]" || ip == "127.0.0.1" {
		return "内网IP"
	}

	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bytes := ghttp.GetBytes(url)
	src := string(bytes)
	srcCharset := "GBK"
	tmp, _ := gcharset.ToUTF8(srcCharset, src)
	json, err := gjson.DecodeToJson(tmp)
	if err != nil {
		return ""
	}
	if json.GetInt("code") == 0 {
		city := json.GetString("city")
		return city
	} else {
		return ""
	}
}

// Capitalize 大写转换
func Capitalize(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}
