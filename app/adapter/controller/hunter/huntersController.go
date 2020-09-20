package hunter

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
	"yu-croco/ddd_on_golang/app/usecase/hunter"
)

type Controller struct{}

func (ctrl Controller) Show(c *gin.Context) {
	hunterId := c.Param("id")
	id, _ := strconv.Atoi(hunterId)
	repo := repositoryImpl.HunterRepositoryImpl{}
	dbResult := repo.FindById(id)

	c.JSON(200, dbResult)
}

// ToDo: メソッド名とActionを統一させる
func (ctrl Controller) Attack(c *gin.Context) {

	monsterId := c.PostForm("monsterId")
	monster, _ := strconv.Atoi(monsterId)
	hunterId := c.Param("id")
	id, _ := strconv.Atoi(hunterId)
	result, err := hunter.AttackMonsterUseCase(id, monster)

	if err.HasErrors() {
		c.JSON(200, err)
	} else {
		c.JSON(200, result)
	}
}
