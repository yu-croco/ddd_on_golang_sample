package monster

import (
	"github.com/gin-gonic/gin"
	helpers2 "yu-croco/ddd_on_golang/pkg/adapter/controller/helpers"
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
	monster2 "yu-croco/ddd_on_golang/pkg/infrastructure/queryImpl/monster"
	repositoryImpl2 "yu-croco/ddd_on_golang/pkg/infrastructure/repositoryImpl"
)

type Controller struct{}

func (ctrl Controller) Show(c *gin.Context) {
	monsterId, monsterIfErr := model2.NewMonsterId(helpers2.ConvertToInt(c.Param("id")))

	if monsterIfErr.HasErrors() {
		helpers2.Response(c, nil, monsterIfErr)
	} else {
		repo := repositoryImpl2.NewMonsterRepositoryImpl()
		result, errs := repo.FindById(*monsterId)

		helpers2.Response(c, result, errs)
	}
}

func (ctrl Controller) Index(c *gin.Context) {
	result := monster2.NewMonsterQueryImpl().FindAll()
	helpers2.Response(c, result, nil)
}
