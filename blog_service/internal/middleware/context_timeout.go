/**
 * @Author: jiangbo
 * @Description:
 * @File:  context_timeout
 * @Version: 1.0.0
 * @Date: 2021/06/20 1:09 下午
 */

package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
