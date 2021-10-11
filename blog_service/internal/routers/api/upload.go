/**
 * @Author: jiangbo
 * @Description:
 * @File:  upload
 * @Version: 1.0.0
 * @Date: 2021/06/20 11:00 上午
 */

package api

import (
	"github.com/gin-gonic/gin"
	"jiangbo.com/blog_service/internal/service"
	"jiangbo.com/blog_service/pkg/app"
	"jiangbo.com/blog_service/pkg/convert"
	"jiangbo.com/blog_service/pkg/errcode"
	"jiangbo.com/blog_service/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}


func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		//global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}

