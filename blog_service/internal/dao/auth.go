/**
 * @Author: jiangbo
 * @Description:
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2021/06/20 1:27 下午
 */

package dao

import "jiangbo.com/blog_service/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
