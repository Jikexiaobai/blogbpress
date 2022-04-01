package api

import (
	"fiber/app/tools/response"
	"fiber/app/tools/utils"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gconv"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"strings"
)

var Thumbnail = new(thumbnailApi)

type thumbnailApi struct{}

func (c *thumbnailApi) LoadRouter(group *ghttp.RouterGroup) {
	group.GET("/*any", c.getImage)
}

func (c *thumbnailApi) getImage(r *ghttp.Request) {
	url := strings.Split(r.URL.Path, "@")
	if len(url) != 2 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(response.CodeMsg(response.PARAM_INVALID)).Send()
	}
	thumbPath := url[0]
	thumbInfo := url[1]
	ext := utils.Ext(thumbPath)
	if ext != ".png" && ext != ".jpeg" && ext != ".jpg" && ext != ".gif" {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(response.CodeMsg(response.PARAM_INVALID)).Send()
	}
	info := strings.Split(thumbInfo, "_")
	if len(info) != 2 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(response.CodeMsg(response.PARAM_INVALID)).Send()
	}
	thumbWidth, thumbHeight := info[0], info[1]

	width := gconv.Int(thumbWidth[1:])
	height := gconv.Int(thumbHeight[1:])
	quality := 50
	// 获取文件名
	fileName := gfile.Name("." + thumbPath)
	dst := "/public/thumb/" + fileName + "/" + thumbWidth + thumbHeight + ext
	//检查文件是否存在
	if gfile.Exists("." + dst) {
		fIn, _ := gfile.Open("." + dst)
		defer fIn.Close()
		origin, fm, err := image.Decode(fIn)
		if err != nil {
			response.Error(r).
				SetCode(response.FAILD).
				SetMessage("打开图片失败").Send()
		}
		switch fm {
		case "jpg":
			jpeg.Encode(r.Response.ResponseWriter, origin, &jpeg.Options{quality})
		case "jpeg":
			jpeg.Encode(r.Response.ResponseWriter, origin, &jpeg.Options{quality})
		case "png":
			png.Encode(r.Response.ResponseWriter, origin)
		case "gif":
			gif.Encode(r.Response.ResponseWriter, origin, &gif.Options{})
		case "bmp":
			bmp.Encode(r.Response.ResponseWriter, origin)
		default:
			response.Error(r).
				SetCode(response.FAILD).
				SetMessage("打开图片失败").Send()
		}
	} else {
		if gfile.Exists("." + "/public/thumb/" + fileName) {
			err := gfile.Remove("." + "/public/thumb/" + fileName)
			if err != nil {
				g.Dump(err.Error())
				response.Error(r).
					SetCode(response.FAILD).
					SetMessage("打开图片失败").Send()
			}
		}
		fIn, _ := gfile.Open("." + thumbPath)
		defer fIn.Close()

		fOut, _ := gfile.Create("." + dst)
		defer fOut.Close()

		origin, fm, err := image.Decode(fIn)
		if err != nil {
			response.Error(r).
				SetCode(response.FAILD).
				SetMessage("打开图片失败").Send()
		}
		if width == 0 || height == 0 {
			width = origin.Bounds().Max.X
			height = origin.Bounds().Max.Y
		}

		canvas := resize.Thumbnail(uint(width), uint(height), origin, resize.Lanczos3)
		switch fm {
		case "jpg":
			header := r.Header
			header.Add("Content-Type", "image/jpg")
			//jpeg.Encode(fOut, thumb)
			jpeg.Encode(fOut, canvas, &jpeg.Options{quality})
			jpeg.Encode(r.Response.ResponseWriter, canvas, &jpeg.Options{quality})
		case "jpeg":
			header := r.Header
			header.Add("Content-Type", "image/jpg")

			//jpeg.Encode(outr.Response.Writer, thumb)
			jpeg.Encode(fOut, canvas, &jpeg.Options{quality})
			jpeg.Encode(r.Response.ResponseWriter, canvas, &jpeg.Options{quality})
		case "png":
			header := r.Header
			header.Add("Content-Type", "image/png")
			png.Encode(fOut, canvas)
			png.Encode(r.Response.ResponseWriter, canvas)
		case "gif":
			header := r.Header
			header.Add("Content-Type", "image/gif")
			gif.Encode(fOut, canvas, &gif.Options{})
			gif.Encode(r.Response.ResponseWriter, canvas, &gif.Options{})
		case "bmp":
			bmp.Encode(fOut, canvas)
		default:
			response.Error(r).
				SetCode(response.FAILD).
				SetMessage("打开图片失败").Send()
		}
	}
}
