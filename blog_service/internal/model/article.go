/**
 * @Author: jiangbo
 * @Description:
 * @File:  article
 * @Version: 1.0.0
 * @Date: 2021/06/19 2:37 下午
 */

package model

import "github.com/jinzhu/gorm"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
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

	db = db.Where("state=?", a.State)
	if err := db.Model(&a).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (a Article) List(db *gorm.DB, offset int, size int) ([]*Article, error) {
	var articles []*Article
	var err error
	if offset >0 && size >0 {
		db = db.Offset(offset).Limit(size)
	}
	if a.Title != "" {
		db = db.Where("title=?", a.Title)
	}
	// 其他字段...
	db = db.Where("state=?", a.State)
	if err = db.Find(&articles).Error; err !=nil{
		return nil, err
	}
	return articles, nil
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

