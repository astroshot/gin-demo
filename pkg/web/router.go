package web

import (
	"astroshot/gintama/pkg/web/controller"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.GET("/hello", controller.Hello)
	Router.GET("/monitor", controller.Monitor)
}
