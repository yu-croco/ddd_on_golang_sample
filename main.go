package main

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/infrastructure"
)

func main() {
	infrastructure.Init()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}