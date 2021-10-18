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

func (d *Dao) CreateArticle(title string, desc string, content string, state uint8, by string, category_id int, tags []int) error {
	var ts []model.Tag

	article := model.Article{
		Title:      title,
		Desc:       desc,
		Content:    content,
		State:      state,
		Model:      &model.Model{CreatedBy: by},
		CategoryId: uint32(category_id),
	}

	if err := article.Create(d.engine); err != nil {
		return err
	}

	for _, v := range tags {
		tmp, _ := d.RetrieveTag(uint32(v))
		ts = append(ts, *tmp)
	}

	d.engine.Model(&article).Association("Tag").Append(ts)
	return nil
}

func (d *Dao) UpdateArticle(id uint32, name string, state uint8, modifiedBy string, content string,
	desc string, category_id int, tags []int) error {
	article := model.Article{
		Model: &model.Model{ID: id},
	}

	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
		"content":     content,
		"desc":        desc,
		"category_id": category_id,
	}

	if err := article.Update(d.engine, values); err != nil {
		return err
	}

	var ts []model.Tag
	for _, v := range tags {
		tmp, _ := d.RetrieveTag(uint32(v))
		ts = append(ts, *tmp)
	}

	d.engine.Model(&article).Association("Tag").Replace(ts)

	return nil
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}
