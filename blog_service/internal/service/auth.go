/**
 * @Author: jiangbo
 * @Description:
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2021/06/20 1:27 下午
 */

package service

import "errors"

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist.")
}
