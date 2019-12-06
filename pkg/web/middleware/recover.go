package middleware

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	view "astroshot/gin-demo/pkg/common/model"
	"astroshot/gin-demo/pkg/util"
)

// Recover recovers server from panic
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				res := view.Fail(-1, util.FailInfo, r)
				c.JSON(http.StatusOK, res)
			}
		}()
		c.Next()
	}
}
