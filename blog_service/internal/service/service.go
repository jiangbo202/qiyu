/**
 * @Author: jiangbo
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2021/06/19 7:08 下午
 */

package service

import (
	"context"
	"jiangbo.com/blog_service/global"
	"jiangbo.com/blog_service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
