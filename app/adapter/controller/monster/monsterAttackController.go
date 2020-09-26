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
	monsterId := model.MonsterId(helpers.ConvertToInt(c.Param("id")))

	result, errs := usecase.AttackHunterUseCase(monsterId, hunter.Id)

	helpers.Response(c, result, errs)
}
