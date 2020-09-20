package hunter

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
	"yu-croco/ddd_on_golang/app/usecase/hunter"
)

type Controller struct{}

func (ctrl Controller) Show(c *gin.Context) {
	hunterId := helpers.ConvertToInt(c.Param("id"))
	repo := repositoryImpl.NewHunterRepositoryImpl()

	dbResult, _ := repo.FindById(hunterId)

	c.JSON(200, dbResult)
}

// ToDo: メソッド名とActionを統一させる
func (ctrl Controller) Attack(c *gin.Context) {
	monsterId := helpers.ConvertToInt(c.PostForm("monsterId"))
	hunterId := helpers.ConvertToInt(c.Param("id"))

	result, err := hunter.AttackMonsterUseCase(hunterId, monsterId)

	if err.HasErrors() {
		c.JSON(200, err)
	} else {
		c.JSON(200, result)
	}
}
