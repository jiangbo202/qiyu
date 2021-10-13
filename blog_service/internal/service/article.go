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
	Title         string `form:"title" binding:"max=20"`
	Desc          string `form:"desc" binding:"max=100" `
	Content       string `form:"content" binding:"max=200"`
	State         uint8  `form:"state,default=1" binding:"required,oneof=0 1"`
}

type ListArticleRequest struct {
	Title         string `form:"title" binding:"max=20"`
	Desc          string `form:"desc" binding:"max=100" `
	Content       string `form:"content" binding:"max=200"`
	State         uint8  `form:"state,default=1" binding:"required,oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,max=20"`
	Desc          string `form:"desc" binding:"max=100" `
	Content       string `form:"content" binding:"max=200"`
	State         uint8  `form:"state,default=1" binding:"required,oneof=0 1"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
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
		param.CreatedBy)
}