package hunter

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/app/domain/model"
	queryImpl "yu-croco/ddd_on_golang/app/infrastructure/queryImpl/hunter"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
)

type HuntersController struct{}

func (ctrl HuntersController) Show(c *gin.Context) {
	hunterId, hunterIdErr := model.NewHunterId(helpers.ConvertToInt(c.Param("id")))

	if hunterIdErr.HasErrors() {
		helpers.Response(c, nil, hunterIdErr)
	} else {
		repo := repositoryImpl.NewHunterRepositoryImpl()
		result, errs := repo.FindById(*hunterId)
		helpers.Response(c, result, errs)
	}
}

func (ctrl HuntersController) Index(c *gin.Context) {
	result := queryImpl.NewHunterQueryImpl().FindAll()
	helpers.Response(c, result, nil)
}
