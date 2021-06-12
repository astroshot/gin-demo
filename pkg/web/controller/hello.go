package controller

import (
	"fmt"
	"net/http"

	"gin-demo/pkg/config"
	"gin-demo/pkg/entity"
	"gin-demo/pkg/util"

	"github.com/gin-gonic/gin"
)

// Hello handles URI /hello
func Hello(c *gin.Context) {
	_ = c.Query("offset")

	conf := config.GetConfInstance()
	path := conf.Server.Path
	fmt.Println(*path)
	res := entity.Success(0, util.SuccessInfo, conf)
	c.JSON(http.StatusOK, res)
}
