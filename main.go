package main

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/adapter/controller/hunter"
	"yu-croco/ddd_on_golang/app/infrastructure"
)

func main() {
	db := infrastructure.Init()
	defer db.Close()

	r := gin.Default()

	hunters := r.Group("/hunters")
	{
		ctrl := hunter.Controller{}
		hunters.GET("/:id", ctrl.Show)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
