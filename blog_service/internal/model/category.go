/**
 * @Author: jiangbo
 * @Description:
 * @File:  category
 * @Version: 1.0.0
 * @Date: 2021/10/13 10:22 下午
 */

package model

import "github.com/jinzhu/gorm"

type Category struct {
	*Model
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func (c Category) TableName() string {
	return "blog_category"
}

func (c Category) Count(db *gorm.DB) (int, error) {
	var count int
	if c.Name != "" {
		db = db.Where("name like ?", "%"+c.Name+"%")
	}
	db = db.Where("status = ?", c.Status)
	if err := db.Model(&c).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (c Category) Retrieve(db *gorm.DB) (*Category, error) {
	category := &Category{}
	err := db.Where("id = ?", c.ID).First(&category).Error
	return category, err
}

func (c Category) List(db *gorm.DB, pageOffset, pageSize int) ([]*Category, error) {
	var categorys []*Category
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if c.Name != "" {
		db = db.Where("name like ?", "%"+c.Name+"%")
	}
	db = db.Where("status=?", c.Status)
	if err = db.Where("is_del=?", 0).Find(&categorys).Error; err != nil {
		return nil, err
	}
	return categorys, nil
}

func (c Category) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}

func (c Category) Update(db *gorm.DB, values interface{}) error {
	//return db.Model(&Tag{}).Where("id=? and is_del=?", t.ID, 0).Update(t).Error
	if err := db.Model(c).Where("id=? and is_del=?", c.ID, 0).Update(values).Error; err != nil {
		return err
	}
	return nil
}

func (c Category) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del=?", c.ID, 0).Delete(&c).Error
}
