package hunter

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/usecase/hunter"
)

type HunterGetMatrialController struct{}

func (ctrl HunterGetMatrialController) Update(c *gin.Context) {
	var monster model.Monster
	c.BindJSON(&monster)
	hunterId, hunterIdErr := model.NewHunterId(helpers.ConvertToInt(c.Param("id")))

	if hunterIdErr.HasErrors() {
		helpers.Response(c, nil, hunterIdErr)
	} else {
		result, errs := hunter.GetMaterialFromMonsterUseCase(*hunterId, monster.Id)
		helpers.Response(c, result, errs)
	}
}
