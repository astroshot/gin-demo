package main

import (
	"flag"
	"fmt"

	"astroshot/gin-demo/pkg/config"
	"astroshot/gin-demo/pkg/web"
)

func main() {
	port := flag.String("port", "8000", "HTTP port")
	env := flag.String("env", "dev", "deploy environment")
	flag.Parse()

	fmt.Println(port)
	fmt.Println(env)
	fmt.Println("Listening to localhost:", *port)
	fmt.Println("Using env:", *env)

	config.GetConfig(env)
	web.Router.Run(":" + *port)
}
