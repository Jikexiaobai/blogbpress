package utils

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gutil"
)

// IsMobile 检测是否使用手机访问
func IsMobile(request *ghttp.Request) bool {
	if via := request.Header.Get("Via"); !gutil.IsEmpty(via) {
		if gstr.Pos(gstr.ToLower(via), "wap") != -1 {
			return true
		}
	}

	if accept := request.Header.Get("Accept"); !gutil.IsEmpty(accept) {
		if gstr.Pos(gstr.ToLower(accept), "vnd.wap.wml") != -1 {
			return true
		}
	}

	if len(request.Header.Values("X-Wap-Profile")) > 0 || len(request.Header.Values("Profile")) > 0 {
		return true
	}

	if userAgent := request.Header.Get("User-Agent"); !gutil.IsEmpty(userAgent) {
		userAgent = gstr.ToLower(userAgent)
		matchStr := gstr.ToLower(`(?i:blackberry|configuration\/cldc|hp |hp-|htc |htc_|htc-|iemobile|kindle|midp|mmp|motorola|mobile|nokia|opera mini|opera |Googlebot-Mobile|YahooSeeker\/M1A1-R2D2|android|iphone|ipod|mobi|palm|palmos|pocket|portalmmm|ppc;|smartphone|sonyericsson|sqh|spv|symbian|treo|up.browser|up.link|vodafone|windows ce|xda |xda_)`)
		// (?i:mobile|ipod|iphone|android|opera mini|blackberry|webos|ucweb|blazer|psp)
		if gregex.IsMatchString(matchStr, userAgent) {
			return true
		}
	}

	return false
}
