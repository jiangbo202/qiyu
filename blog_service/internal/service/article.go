/**
 * @Author: jiangbo
 * @Description:
 * @File:  article
 * @Version: 1.0.0
 * @Date: 2021/06/20 5:21 下午
 */

package service

import (
	"jiangbo.com/blog_service/internal/model"
	"jiangbo.com/blog_service/pkg/app"
)

type CountArticleRequest struct {
	Title   string `json:"title" binding:"max=20"`
	Desc    string `json:"desc" binding:"max=100" `
	Content string `json:"content" binding:"max=200"`
	State   uint8  `json:"state,default=1" binding:"required,oneof=0 1"`
}

type ListArticleRequest struct {
	Title   string `json:"title" binding:"max=20"`
	Desc    string `json:"desc" binding:"max=100" `
	Content string `json:"content" binding:"max=200"`
	State   uint8  `json:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title      string `json:"title" binding:"required,max=20"`
	Desc       string `json:"desc" binding:"max=100" `
	Content    string `json:"content" binding:"max=200"`
	State      uint8  `json:"state,default=1" binding:"required,oneof=0 1"`
	CreatedBy  string `json:"created_by" binding:"required,min=3,max=100"`
	CategoryId int `json:"category_id" binding:"required"`
	Tags       []int `json:"tags"`
}

type UpdateArticleRequest struct {
	ID         uint32 `json:"id" binding:"required,gte=1"`
	ModifiedBy string `json:"modified_by" binding:"required,min=3,max=100"`
	Title      string `json:"title" binding:"required,max=20"`
	Desc       string `json:"desc" binding:"max=100" `
	Content    string `json:"content" binding:"max=200"`
	State      uint8  `json:"state,default=1" binding:"required,oneof=0 1"`
	CategoryId int `json:"category_id" binding:"required"`
	Tags       []int `json:"tags"`
}

type DeleteArticleRequest struct {
	ID uint32 `json:"id" binding:"required,gte=1"`
}

func (svc *Service) CountArticle(param *CountArticleRequest) (int, error) {
	return svc.dao.CountArticle(param.Title, param.Desc, param.Content, param.State)
}

func (svc *Service) GetArticleList(param *ListArticleRequest, pager *app.Pager) ([]*model.Article, error) {
	return svc.dao.GetArticleList(param.Title, param.Desc, param.Content,
		param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title, param.Desc, param.Content, param.State,
		param.CreatedBy, param.CategoryId, param.Tags)
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.State, param.ModifiedBy, param.Content, param.Desc,
		param.CategoryId, param.Tags)
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}