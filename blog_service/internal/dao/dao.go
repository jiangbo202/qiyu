/**
 * @Author: jiangbo
 * @Description:
 * @File:  dao
 * @Version: 1.0.0
 * @Date: 2021/06/19 7:03 下午
 */

package dao

import (
	"github.com/jinzhu/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
