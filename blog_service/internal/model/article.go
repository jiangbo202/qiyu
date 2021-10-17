/**
 * @Author: jiangbo
 * @Description:
 * @File:  article
 * @Version: 1.0.0
 * @Date: 2021/06/19 2:37 下午
 */

package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title      string   `json:"title"`
	Desc       string   `json:"desc"`
	Content    string   `json:"content"`
	State      uint8    `json:"state"`
	CategoryId uint32   `json:"category_id"`
	Category   Category `json:"category" gorm:"foreignkey:CategoryID"` // 关联外键
	Tag        []Tag    `json:"tag" gorm:"many2many:blog_article_tag"` // 多对多
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Count(db *gorm.DB) (int, error) {
	var count int
	if a.Title != "" {
		db = db.Where("title=?", a.Title)
	}
	// 其他字段...

	//db = db.Where("state=?", a.State)
	if err := db.Model(&a).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (a Article) List(db *gorm.DB, offset int, size int) ([]*Article, error) {
	var articles []*Article
	var err error
	if offset > 0 && size > 0 {
		db = db.Offset(offset).Limit(size)
	}
	if a.Title != "" {
		db = db.Where("title=?", a.Title)
	}
	// 其他字段...
	//db = db.Where("state=?", a.State)
	fmt.Println(a)
	if err = db.Preload("Category").Preload("Tag").Find(&articles).Error; err != nil {
		return nil, err
	}
	fmt.Println(articles)
	return articles, nil
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	//return db.Model(&Tag{}).Where("id=? and is_del=?", t.ID, 0).Update(t).Error
	if err := db.Model(a).Where("id=? and is_del=?", a.ID, 0).Update(values).Error; err != nil {
		return err
	}
	return nil
}