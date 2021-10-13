/**
 * @Author: jiangbo
 * @Description:
 * @File:  category
 * @Version: 1.0.0
 * @Date: 2021/10/13 10:22 下午
 */

package model

type Category struct {
	*Model
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func (a Category) TableName() string {
	return "blog_category"
}

