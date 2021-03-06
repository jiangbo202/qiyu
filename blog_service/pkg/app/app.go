/**
 * @Author: jiangbo
 * @Description:
 * @File:  app
 * @Version: 1.0.0
 * @Date: 2021/06/19 5:15 下午
 */

package app

import (
	"github.com/gin-gonic/gin"
	"jiangbo.com/blog_service/pkg/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r Response) ToResponse(data interface{}) {
	//if data == nil {
	//	data = gin.H{}
	//}
	data = gin.H{
		"code": 0,
		"msg": "操作成功",
		"data": data,
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r Response) ToResponseList(list interface{}, totalRows int)  {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page: GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details)>0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}