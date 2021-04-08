package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	view "gin-demo/pkg/common/model"
	"gin-demo/pkg/util"
)

func MethodNotSupportedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := view.FailWithDefaultCode(util.MethodNotSupported, nil)
		c.JSON(http.StatusOK, res)
		c.Next()
	}
}

func URINotFoundMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := view.FailWithDefaultCode(util.URINotFound, nil)
		c.JSON(http.StatusOK, res)
		c.Next()
	}
}
