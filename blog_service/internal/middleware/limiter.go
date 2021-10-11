/**
 * @Author: jiangbo
 * @Description:
 * @File:  limiter
 * @Version: 1.0.0
 * @Date: 2021/06/20 1:07 下午
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"jiangbo.com/blog_service/pkg/app"
	"jiangbo.com/blog_service/pkg/errcode"
	"jiangbo.com/blog_service/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
