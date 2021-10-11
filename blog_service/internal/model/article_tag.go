/**
 * @Author: jiangbo
 * @Description:
 * @File:  article_tag
 * @Version: 1.0.0
 * @Date: 2021/06/19 2:39 下午
 */

package model

type ArticleTag struct {
	*Model
	TagID uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (t ArticleTag) TableName() string {
	return "blog_article_tag"
}
