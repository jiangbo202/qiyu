/**
 * @Author: jiangbo
 * @Description:
 * @File:  recover
 * @Version: 1.0.0
 * @Date: 2021/06/20 12:58 下午
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"jiangbo.com/blog_service/global"
	"jiangbo.com/blog_service/pkg/app"
	"jiangbo.com/blog_service/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
