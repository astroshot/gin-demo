package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Monitor returns ok for echo server health
func Monitor(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
