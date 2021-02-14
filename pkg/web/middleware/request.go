package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	view "gin-demo/pkg/common/model"
	"gin-demo/pkg/util"
)

func NoMethodMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := view.Fail(-1, util.MethodNotFound, nil)
		c.JSON(http.StatusOK, res)
		c.Next()
	}
}
