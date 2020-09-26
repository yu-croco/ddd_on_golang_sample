package monster

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/app/domain/model"
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
		result, errs := usecase.AttackHunterUseCase(*monsterId, hunter.Id)
		helpers.Response(c, result, errs)
	}

}
