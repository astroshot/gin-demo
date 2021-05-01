package main

import (
	_ "encoding/json"
	"flag"
	"fmt"

	"gin-demo/pkg/config"
	"gin-demo/pkg/service"
	"gin-demo/pkg/web"
)

var (
	Version   string
	GitCommit string
	GoVersion string
	BuildTime string
)

func fullVersion() string {
	return fmt.Sprintf("Version: %6s, Git commit: %6s, Go version: %6s, Build time: %6s", Version, GitCommit, GoVersion, BuildTime)
}

func main() {
	env := flag.String("env", "dev", "deploy environment")
	showVersion := flag.Bool("v", false, "Current Version")
	flag.Parse()

	if *showVersion {
		fmt.Println(fullVersion())
		return
	}

	fmt.Println(env)
	fmt.Println("Using env:", *env)

	conf := config.GetConfig(env)
	// confStr, _ := json.Marshal(conf)
	// fmt.Println(string(confStr))
	service.InitService()
	web.ConfigLog(conf)
	web.MapURI(conf)
	web.Router.Run(":" + *conf.Server.Port)
}
