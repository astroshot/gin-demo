package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetQueryStr(c *gin.Context, key string) *string {
	if val, ok := c.GetQuery(key); ok {
		return &val
	}
	return nil
}

func GetQueryInt(c *gin.Context, key string) *int {
	val := GetQueryStr(c, key)
	if val == nil {
		return nil
	}

	if intVal, err := strconv.Atoi(*val); err == nil {
		return &intVal
	}

	return nil
}
