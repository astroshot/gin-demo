package controller

import (
	"gin-demo/pkg/config"

	"fmt"

	"github.com/gin-gonic/gin"
	"net/http"
)

// Hello handles URI /hello
func Hello(c *gin.Context) {
	_ = c.Query("offset")

	conf := config.GetConfig(config.GetEnv())
	fmt.Println(conf.Server.Path)
	c.String(http.StatusOK, *conf.Server.Path)
}
