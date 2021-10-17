/**
 * @Author: jiangbo
 * @Description:
 * @File:  category
 * @Version: 1.0.0
 * @Date: 2021/10/17 5:02 下午
 */

package service

import (
	"jiangbo.com/blog_service/internal/model"
	"jiangbo.com/blog_service/pkg/app"
)

type CountCategoryRequest struct {
	Name  string `json:"name" binding:"max=100"`
	Status int `json:"status,default=1" binding:"oneof=0 1"`
}

type RetrieveCategoryRequest struct {
	ID         uint32 `json:"id" binding:"required,gte=1"`
}

type CategoryListRequest struct {
	Name  string `json:"name" binding:"max=100"`
	Status int  `json:"status,default=1" binding:"oneof=0 1"`
}

type CreateCategoryRequest struct {
	Name      string `json:"name" binding:"required,min=3,max=100"`
	CreatedBy string `json:"created_by" binding:"required,min=3,max=100"`
	Status     int  `json:"status,default=1" binding:"oneof=0 1"`
}

type UpdateCategoryRequest struct {
	ID         uint32 `json:"id" binding:"required,gte=1"`
	Name       string `json:"name" binding:"min=3,max=100"`
	Status      uint8  `json:"status" binding:"oneof=0 1"`
	ModifiedBy string `json:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteCategoryRequest struct {
	ID uint32 `json:"id" binding:"required,gte=1"`
}

func (svc *Service) CountCategory(param *CountCategoryRequest) (int, error) {
	return svc.dao.CountCategory(param.Name, param.Status)
}

func (svc *Service) RetrieveCategory(param *RetrieveCategoryRequest) (*model.Category, error) {
	return svc.dao.RetrieveCategory(param.ID)
}

func (svc *Service) GetCategoryList(param *CategoryListRequest, pager *app.Pager) ([]*model.Category, error) {
	return svc.dao.GetCategoryList(param.Name, param.Status, pager.Page, pager.PageSize)
}

func (svc *Service) CreateCategory(param *CreateCategoryRequest) error {
	return svc.dao.CreateCategory(param.Name, param.Status, param.CreatedBy)
}

func (svc *Service) UpdateCategory(param *UpdateCategoryRequest) error {
	return svc.dao.UpdateCategory(param.ID, param.Name, param.Status, param.ModifiedBy)
}

func (svc *Service) DeleteCategory(param *DeleteCategoryRequest) error {
	return svc.dao.DeleteCategory(param.ID)
}


