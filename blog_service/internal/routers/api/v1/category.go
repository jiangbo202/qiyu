/**
 * @Author: jiangbo
 * @Description:
 * @File:  category
 * @Version: 1.0.0
 * @Date: 2021/10/17 4:58 下午
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"jiangbo.com/blog_service/global"
	"jiangbo.com/blog_service/internal/service"
	"jiangbo.com/blog_service/pkg/app"
	"jiangbo.com/blog_service/pkg/convert"
	"jiangbo.com/blog_service/pkg/errcode"
)

type Category struct{}

func NewCategory() Category {
	return Category{}
}

//返回数据，数据转换
//type CategoryRes struct {
//	Id         int    `json:"id"`
//	Name       string `json:"name"`
//	State      uint8  `json:"state"`
//	CreatedBy  string `json:"created_by"`
//	ModifiedBy string `json:"modified_by"`
//	CreatedAt  string `json:"created_at"`
//	ModifiedAt string `json:"modified_at"`
//	DeletedAt  string `json:"deleted_at"`
//	IsDel      uint8  `json:"is_del"`
//}

func (ct Category) Get(c *gin.Context) {
	param := service.RetrieveCategoryRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	category, err := svc.RetrieveCategory(&param)
	if err != nil {
		global.Logger.Errorf("svc.RetrieveCategory err: %v", err)
		response.ToErrorResponse(errcode.ErrorRetrieveCategoryFail)
		return
	}
	response.ToResponse(category)
	return
}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.CategorySwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (ct Category) List(c *gin.Context) {
	param := service.CategoryListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountCategory(&service.CountCategoryRequest{Name: param.Name, Status: param.Status})
	if err != nil {
		global.Logger.Errorf("svc.CountCategory err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountCategoryFail)
		return
	}

	tags, err := svc.GetCategoryList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetCategoryList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetCategoryListFail)
		return
	}

	//tagList := make([]CategoryRes, len(tags))
	//for k, v := range tags {
	//	fmt.Println(v.DeletedAt)
	//	_tag := CategoryRes{
	//		Id: cast.ToInt(v.ID),
	//		Name: v.Name,
	//		State: v.State,
	//		CreatedBy: v.CreatedBy,
	//		ModifiedBy: v.ModifiedBy,
	//		CreatedAt: time_parse.TimeCSTLayoutString(v.CreatedAt),
	//		ModifiedAt: time_parse.TimeCSTLayoutString(v.ModifiedAt),
	//		DeletedAt: time_parse.TimeCSTLayoutString(v.DeletedAt),
	//		IsDel: v.IsDel,
	//	}
	//	tagList[k] = _tag
	//}
	response.ToResponseList(tags, totalRows)
	return
}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Category "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (ct Category) Create(c *gin.Context) {
	param := service.CreateCategoryRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateCategory(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateCategory err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateCategoryFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "类别 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Category "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (ct Category) Update(c *gin.Context) {
	param := service.UpdateCategoryRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateCategory(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateCategory err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateCategoryFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "类别 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (ct Category) Delete(c *gin.Context) {
	param := service.DeleteCategoryRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteCategory(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteCategory err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteCategoryFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
