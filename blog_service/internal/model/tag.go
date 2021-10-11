/**
 * @Author: jiangbo
 * @Description:
 * @File:  tag
 * @Version: 1.0.0
 * @Date: 2021/06/19 2:35 下午
 */

package model

import (
	"github.com/jinzhu/gorm"
	"jiangbo.com/blog_service/pkg/app"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

//	user := &v1.User{}
//	err := u.db.Where("name = ?", username).First(&user).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, errors.WithCode(code.ErrUserNotFound, err.Error())
//		}
//
//		return nil, errors.WithCode(code.ErrDatabase, err.Error())
//	}
//
//	return user, nil


func (t Tag) Retrieve(db *gorm.DB) (*Tag, error) {
	tag := &Tag{}
	err := db.Where("id = ?", t.ID).First(&tag).Error
	return tag, err
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name=?", t.Name)
	}
	db = db.Where("state=?", t.State)
	if err = db.Where("is_del=?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	//return db.Model(&Tag{}).Where("id=? and is_del=?", t.ID, 0).Update(t).Error
	if err := db.Model(t).Where("id=? and is_del=?", t.ID, 0).Update(values).Error; err != nil {
		return err
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del=?", t.ID, 0).Delete(&t).Error
}
