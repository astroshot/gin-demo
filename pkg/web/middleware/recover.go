package middleware

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	entity "gin-demo/pkg/entity"
	"gin-demo/pkg/util"
)

// Recover recovers server from panic
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				res := entity.Fail(-1, util.FailInfo, r)
				c.JSON(http.StatusOK, res)
			}
		}()
		c.Next()
	}
}
