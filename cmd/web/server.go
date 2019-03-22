package main

import (
	"astroshot/gintama/pkg/web"
	"flag"
	"fmt"
)

func main() {
	port := flag.String("port", "8000", "HTTP port")
	flag.Parse()
	fmt.Println("Listening to localhost:", *port)
	web.Router.Run(":" + *port)
}
