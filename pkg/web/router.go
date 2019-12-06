package web

import (
	"astroshot/gin-demo/pkg/config"
	"astroshot/gin-demo/pkg/web/controller"

	"github.com/gin-gonic/gin"
)

// Router maps URI to handle function
var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.Use(GetLogger())
	Router.Use(gin.Recovery())

}

func MapURI(conf *config.Config) {
	v := Router.Group(*conf.Server.Path)
	v.GET("/hello", controller.Hello)
	v.GET("/monitor", controller.Monitor)

	v.GET("/users/:token", controller.GetUserByID)
	v.GET("/users", controller.ListUsers)
	v.POST("/users", controller.AddUser)
}
