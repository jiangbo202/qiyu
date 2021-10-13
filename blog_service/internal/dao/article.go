/**
 * @Author: jiangbo
 * @Description:
 * @File:  article
 * @Version: 1.0.0
 * @Date: 2021/06/20 5:20 下午
 */

package dao

import (
	"jiangbo.com/blog_service/internal/model"
	"jiangbo.com/blog_service/pkg/app"
)

func (d *Dao) CountArticle(title string, desc string, content string, state uint8) (int, error) {
	article := model.Article{Title: title, Desc: desc, Content: content, State: state}
	return article.Count(d.engine)
}

func (d *Dao) GetArticleList(title string, desc string, content string, state uint8, page int, pageSize int) ([]*model.Article, error) {
	article := model.Article{Title: title, Desc: desc, Content: content, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateArticle(title string, desc string, content string, state uint8, by string) error {
	article := model.Article{
		Title: title,
		Desc: desc,
		Content: content,
		State: state,
		Model: &model.Model{CreatedBy: by},
	}
	return article.Create(d.engine)
}
