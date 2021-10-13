/**
 * @Author: jiangbo
 * @Description:
 * @File:  tag
 * @Version: 1.0.0
 * @Date: 2021/06/19 2:49 下午
 */

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"jiangbo.com/blog_service/global"
	"jiangbo.com/blog_service/internal/service"
	"jiangbo.com/blog_service/pkg/app"
	"jiangbo.com/blog_service/pkg/convert"
	"jiangbo.com/blog_service/pkg/errcode"
	"jiangbo.com/blog_service/pkg/time_parse"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

type TagRes struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	State      uint8  `json:"state"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedAt  string `json:"created_at"`
	ModifiedAt string `json:"modified_at"`
	DeletedAt  string `json:"deleted_at"`
	IsDel      uint8  `json:"is_del"`
}

func (t Tag) Get(c *gin.Context) {
	param := service.RetrieveTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	tag, err := svc.RetrieveTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.RetrieveTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorRetrieveTagFail)
		return
	}
	tagRes := TagRes{
		Id: cast.ToInt(tag.ID),
		Name: tag.Name,
		State: tag.State,
		CreatedBy: tag.CreatedBy,
		ModifiedBy: tag.ModifiedBy,
		CreatedAt: time_parse.TimeCSTLayoutString(tag.CreatedAt),
		ModifiedAt: time_parse.TimeCSTLayoutString(tag.ModifiedAt),
		DeletedAt: time_parse.TimeCSTLayoutString(tag.DeletedAt),
		IsDel: tag.IsDel,
	}
	response.ToResponse(tagRes)
	return
}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	tagList := make([]TagRes, len(tags))
	for k, v := range tags {
		fmt.Println(v.DeletedAt)
		_tag := TagRes{
			Id: cast.ToInt(v.ID),
			Name: v.Name,
			State: v.State,
			CreatedBy: v.CreatedBy,
			ModifiedBy: v.ModifiedBy,
			CreatedAt: time_parse.TimeCSTLayoutString(v.CreatedAt),
			ModifiedAt: time_parse.TimeCSTLayoutString(v.ModifiedAt),
			DeletedAt: time_parse.TimeCSTLayoutString(v.DeletedAt),
			IsDel: v.IsDel,
		}
		tagList[k] = _tag
	}

	response.ToResponseList(tagList, totalRows)
	return
}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
