package main

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/adapter/controller/hunter"
	"yu-croco/ddd_on_golang/app/adapter/controller/monster"
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
		hunters.GET("/", ctrl.Index)
		hunters.PUT("/:id/attack", ctrl.Attack)
	}
	monsters := r.Group("/monsters")
	{
		ctrl := monster.Controller{}
		monsters.GET("/:id", ctrl.Show)
		monsters.GET("/", ctrl.Index)
	}

	r.Run()
}
