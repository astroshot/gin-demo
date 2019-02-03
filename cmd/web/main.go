package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("port", "8000", "HTTP port")
	flag.Parse()
	fmt.Println("Listening to localhost:%s", *port)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})
	r.Run(":" + *port)
}
