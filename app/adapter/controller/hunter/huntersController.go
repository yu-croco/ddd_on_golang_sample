package hunter

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/adapter/controller/helpers"
	hunter2 "yu-croco/ddd_on_golang/app/infrastructure/queryImpl/hunter"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
	"yu-croco/ddd_on_golang/app/usecase/hunter"
)

type Controller struct{}

func (ctrl Controller) Show(c *gin.Context) {
	hunterId := helpers.ConvertToInt(c.Param("id"))
	repo := repositoryImpl.NewHunterRepositoryImpl()

	dbResult, err := repo.FindById(hunterId)

	helpers.Response(c, dbResult, err)
}

// ToDo: メソッド名とActionを統一させる
func (ctrl Controller) Attack(c *gin.Context) {
	monsterId := helpers.ConvertToInt(c.PostForm("monsterId"))
	hunterId := helpers.ConvertToInt(c.Param("id"))

	result, err := hunter.AttackMonsterUseCase(hunterId, monsterId)

	helpers.Response(c, result, err)
}

func (ctrl Controller) Index(c *gin.Context) {
	result := hunter2.NewHunterQueryImpl().FindAll()
	helpers.Response(c, result, nil)
}

// ToDo: メソッド名とActionを統一させる
func (ctrl Controller) GetMaterial(c *gin.Context) {
	monsterId := helpers.ConvertToInt(c.PostForm("monsterId"))
	hunterId := helpers.ConvertToInt(c.Param("id"))

	result, err := hunter.GetMaterialFromMonsterUseCase(hunterId, monsterId)

	helpers.Response(c, result, err)
}
