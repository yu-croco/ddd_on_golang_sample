package hunter

import (
	"github.com/gin-gonic/gin"
	helpers2 "yu-croco/ddd_on_golang/pkg/adapter/controller/helpers"
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
	repositoryImpl2 "yu-croco/ddd_on_golang/pkg/infrastructure/repositoryImpl"
	"yu-croco/ddd_on_golang/pkg/usecase/hunter"
)

type HunterGetMatrialController struct{}

func (ctrl HunterGetMatrialController) Update(c *gin.Context) {
	var monster model2.Monster
	c.BindJSON(&monster)
	hunterId, hunterIdErr := model2.NewHunterId(helpers2.ConvertToInt(c.Param("id")))

	if hunterIdErr.HasErrors() {
		helpers2.Response(c, nil, hunterIdErr)
	} else {
		hunterRepository := repositoryImpl2.NewHunterRepositoryImpl()
		monsterRepository := repositoryImpl2.NewMonsterRepositoryImpl()

		result, errs := hunter.NewGetMaterialFromMonsterUseCaseImpl(*hunterId, monster.Id, hunterRepository, monsterRepository).
			Exec()
		helpers2.Response(c, result, errs)
	}
}
