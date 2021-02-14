package middleware

import (
	"github.com/gin-gonic/gin"

	"gin-demo/pkg/helper"
	"gin-demo/pkg/util"
)

// TraceMiddleware 跟踪ID中间件
func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 优先从请求头中获取请求ID，如果没有则使用UUID
		traceID := c.GetHeader("X-Request-Id")
		if traceID == "" {
			traceID = helper.GetTraceID()
		}
		c.Set(util.TraceIDKey, traceID)
		c.Next()
	}
}
