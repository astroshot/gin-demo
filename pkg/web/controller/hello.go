package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	_ = c.Query("offset")
	c.String(http.StatusOK, "Hello!")
}
