package hunter

import (
	"github.com/gin-gonic/gin"
	helpers2 "yu-croco/ddd_on_golang/pkg/adapter/controller/helpers"
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
	"yu-croco/ddd_on_golang/pkg/infrastructure/queryImpl/hunter"
	repositoryImpl2 "yu-croco/ddd_on_golang/pkg/infrastructure/repositoryImpl"
)

type HuntersController struct{}

func (ctrl HuntersController) Show(c *gin.Context) {
	hunterId, hunterIdErr := model2.NewHunterId(helpers2.ConvertToInt(c.Param("id")))

	if hunterIdErr.HasErrors() {
		helpers2.Response(c, nil, hunterIdErr)
	} else {
		repo := repositoryImpl2.NewHunterRepositoryImpl()
		result, errs := repo.FindById(*hunterId)
		helpers2.Response(c, result, errs)
	}
}

func (ctrl HuntersController) Index(c *gin.Context) {
	result := hunter.NewHunterQueryImpl().FindAll()
	helpers2.Response(c, result, nil)
}
