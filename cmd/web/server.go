package main

import (
	"encoding/json"
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
	confStr, _ := json.Marshal(conf)
	fmt.Println(string(confStr))
	service.InitService()
	web.ConfigLog(conf)
	web.MapURI(conf)
	web.Router.Run(":" + *conf.Server.Port)
}
