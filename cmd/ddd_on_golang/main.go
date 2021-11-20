package main

import (
	"github.com/gin-gonic/gin"
	hunter2 "yu-croco/ddd_on_golang/pkg/adapter/controller/hunter"
	monster2 "yu-croco/ddd_on_golang/pkg/adapter/controller/monster"
	infrastructure2 "yu-croco/ddd_on_golang/pkg/infrastructure"
)

func main() {
	db := infrastructure2.Init()
	defer db.Close()

	r := gin.Default()

	hunters := r.Group("/hunters")
	{
		huntersCtrl := hunter2.HuntersController{}
		hunters.GET("/:id", huntersCtrl.Show)
		hunters.GET("/", huntersCtrl.Index)

		hunterAttackCtrl := hunter2.HunterAttackController{}
		hunters.PUT("/:id/attack", hunterAttackCtrl.Update)

		hunterGetMaterialCtrl := hunter2.HunterGetMatrialController{}
		hunters.POST("/:id/get_material_from_monster", hunterGetMaterialCtrl.Update)
	}

	monsters := r.Group("/monsters")
	{
		monsterCtrl := monster2.Controller{}
		monsters.GET("/:id", monsterCtrl.Show)
		monsters.GET("/", monsterCtrl.Index)

		monsterAttackCtrl := monster2.MonsterAttackController{}
		monsters.PUT("/:id/attack", monsterAttackCtrl.Update)
	}

	r.Run()
}
