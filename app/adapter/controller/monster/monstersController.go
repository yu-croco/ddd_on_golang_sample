package monster

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
)

type Controller struct{}

func (ctrl Controller) Show(c *gin.Context) {
	hunterId := helpers.ConvertToInt(c.Param("id"))
	repo := repositoryImpl.NewMonsterRepositoryImpl()

	dbResult, err := repo.FindById(hunterId)
	helpers.Response(c, dbResult, err)
}
