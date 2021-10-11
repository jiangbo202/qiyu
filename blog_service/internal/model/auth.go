/**
 * @Author: jiangbo
 * @Description:
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2021/06/20 1:20 下午
 */

package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*Model
	AppKey string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}