package monster

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure/queryImpl/monster"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
)

type Controller struct{}

func (ctrl Controller) Show(c *gin.Context) {
	monsterId, monsterIfErr := model.NewMonsterId(helpers.ConvertToInt(c.Param("id")))

	if monsterIfErr.HasErrors() {
		helpers.Response(c, nil, monsterIfErr)
	} else {
		repo := repositoryImpl.NewMonsterRepositoryImpl()
		result, errs := repo.FindById(*monsterId)

		helpers.Response(c, result, errs)
	}
}

func (ctrl Controller) Index(c *gin.Context) {
	result := monster.NewMonsterQueryImpl().FindAll()
	helpers.Response(c, result, nil)
}
