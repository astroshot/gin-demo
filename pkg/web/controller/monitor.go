package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Monitor(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
