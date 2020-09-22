package monster

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/app/infrastructure/queryImpl/monster"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
)

type Controller struct{}

func (ctrl Controller) Show(c *gin.Context) {
	monsterId := helpers.ConvertToInt(c.Param("id"))

	repo := repositoryImpl.NewMonsterRepositoryImpl()
	result, errs := repo.FindById(monsterId)

	helpers.Response(c, result, errs)
}

func (ctrl Controller) Index(c *gin.Context) {
	result := monster.NewMonsterQueryImpl().FindAll()
	helpers.Response(c, result, nil)
}
