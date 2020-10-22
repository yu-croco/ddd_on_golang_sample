package monster

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
	usecase "yu-croco/ddd_on_golang/app/usecase/monster"
)

type MonsterAttackController struct{}

func (ctrl MonsterAttackController) Update(c *gin.Context) {
	var hunter model.Hunter
	c.BindJSON(&hunter)
	monsterId, monsterIdErr := model.NewMonsterId(helpers.ConvertToInt(c.Param("id")))
	if monsterIdErr.HasErrors() {
		helpers.Response(c, nil, monsterIdErr)
	} else {
		hunterRepository := repositoryImpl.NewHunterRepositoryImpl()
		monsterRepository := repositoryImpl.NewMonsterRepositoryImpl()

		result, errs := usecase.AttackHunterUseCase(*monsterId, hunter.Id, hunterRepository, monsterRepository)
		helpers.Response(c, result, errs)
	}

}
