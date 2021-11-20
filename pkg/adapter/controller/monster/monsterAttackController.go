package monster

import (
	"github.com/gin-gonic/gin"
	helpers2 "yu-croco/ddd_on_golang/pkg/adapter/controller/helpers"
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
	repositoryImpl2 "yu-croco/ddd_on_golang/pkg/infrastructure/repositoryImpl"
	"yu-croco/ddd_on_golang/pkg/usecase/monster"
)

type MonsterAttackController struct{}

func (ctrl MonsterAttackController) Update(c *gin.Context) {
	var hunter model2.Hunter
	c.BindJSON(&hunter)
	monsterId, monsterIdErr := model2.NewMonsterId(helpers2.ConvertToInt(c.Param("id")))
	if monsterIdErr.HasErrors() {
		helpers2.Response(c, nil, monsterIdErr)
	} else {
		hunterRepository := repositoryImpl2.NewHunterRepositoryImpl()
		monsterRepository := repositoryImpl2.NewMonsterRepositoryImpl()

		result, errs := monster.NewAttackHunterUseCaseImpl(hunter.Id, *monsterId, hunterRepository, monsterRepository).
			Exec()
		helpers2.Response(c, result, errs)
	}

}
