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
		huntersCtrl := hunter.HuntersController{}
		hunters.GET("/:id", huntersCtrl.Show)
		hunters.GET("/", huntersCtrl.Index)

		hunterAttackCtrl := hunter.HunterAttackController{}
		hunters.PUT("/:id/attack", hunterAttackCtrl.Update)

		hunterGetMaterialCtrl := hunter.HunterGetMatrialController{}
		hunters.POST("/:id/get_material_from_monster", hunterGetMaterialCtrl.Update)
	}

	monsters := r.Group("/monsters")
	{
		monsterCtrl := monster.Controller{}
		monsters.GET("/:id", monsterCtrl.Show)
		monsters.GET("/", monsterCtrl.Index)

		monsterAttackCtrl := monster.MonsterAttackController{}
		monsters.PUT("/:id/attack", monsterAttackCtrl.Update)
	}

	r.Run()
}
