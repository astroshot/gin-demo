package main

import (
	"flag"
	"fmt"

	"astroshot/gin-demo/pkg/config"
	"astroshot/gin-demo/pkg/service"
	"astroshot/gin-demo/pkg/web"
)

func main() {
	env := flag.String("env", "dev", "deploy environment")
	flag.Parse()

	fmt.Println(env)
	fmt.Println("Using env:", *env)

	conf := config.GetConfig(env)
	service.InitService()
	web.Router.Run(":" + *conf.Server.Port)
}
