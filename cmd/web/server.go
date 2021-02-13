package main

import (
	"flag"
	"fmt"

	"gin-demo/pkg/config"
	"gin-demo/pkg/service"
	"gin-demo/pkg/web"
)

func main() {
	env := flag.String("env", "dev", "deploy environment")
	flag.Parse()

	fmt.Println(env)
	fmt.Println("Using env:", *env)

	conf := config.GetConfig(env)
	service.InitService()
	web.MapURI(conf)
	web.Router.Run(":" + *conf.Server.Port)
}
