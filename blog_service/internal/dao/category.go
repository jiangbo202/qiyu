/**
 * @Author: jiangbo
 * @Description:
 * @File:  category
 * @Version: 1.0.0
 * @Date: 2021/10/17 4:53 下午
 */

package dao

import (
	"jiangbo.com/blog_service/internal/model"
	"jiangbo.com/blog_service/pkg/app"
)

func (d *Dao) CountCategory(name string, status int) (int, error) {
	tag := model.Category{Name: name, Status: status}
	return tag.Count(d.engine)
}

func (d *Dao) RetrieveCategory(id uint32) (*model.Category, error) {
	tag := model.Category{
		Model: &model.Model{ID: id},
	}
	return tag.Retrieve(d.engine)
}

func (d *Dao) GetCategoryList(name string, status int, page, pageSize int) ([]*model.Category, error) {
	tag := model.Category{Name: name, Status: status}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateCategory(name string, status int, createdBy string) error {
	tag := model.Category{
		Name:   name,
		Status: status,
		Model:  &model.Model{CreatedBy: createdBy},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateCategory(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Category{
		Model: &model.Model{ID: id},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}
	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteCategory(id uint32) error {
	tag := model.Category{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}
