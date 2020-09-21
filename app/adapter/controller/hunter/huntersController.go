package hunter

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/app/domain/model"
	queryImpl "yu-croco/ddd_on_golang/app/infrastructure/queryImpl/hunter"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
	"yu-croco/ddd_on_golang/app/usecase/hunter"
)

type Controller struct{}

func (ctrl Controller) Show(c *gin.Context) {
	hunterId, hunterIdErr := model.NewHunterId(helpers.ConvertToInt(c.Param("id")))
	if hunterIdErr.HasErrors() {
		helpers.Response(c, nil, hunterIdErr)
	} else {
		repo := repositoryImpl.NewHunterRepositoryImpl()
		dbResult, err := repo.FindById(hunterId)

		helpers.Response(c, dbResult, err)
	}
}

// ToDo: メソッド名とActionを統一させる
func (ctrl Controller) Attack(c *gin.Context) {
	monsterId, monsterIdErr := model.NewMonsterId(helpers.ConvertToInt(c.PostForm("monsterId")))
	hunterId, hunterIdErr := model.NewHunterId(helpers.ConvertToInt(c.Param("id")))

	if monsterIdErr.HasErrors() || hunterIdErr.HasErrors() {
		errs := monsterIdErr.Concat(hunterIdErr)
		helpers.Response(c, nil, &errs)
	} else {
		result, err := hunter.AttackMonsterUseCase(hunterId, monsterId)

		helpers.Response(c, result, err)
	}
}

func (ctrl Controller) Index(c *gin.Context) {
	result := queryImpl.NewHunterQueryImpl().FindAll()
	helpers.Response(c, result, nil)
}

// ToDo: メソッド名とActionを統一させる
func (ctrl Controller) GetMaterial(c *gin.Context) {
	monsterId, monsterIdErr := model.NewMonsterId(helpers.ConvertToInt(c.PostForm("monsterId")))
	hunterId, hunterIdErr := model.NewHunterId(helpers.ConvertToInt(c.Param("id")))

	if monsterIdErr.HasErrors() || hunterIdErr.HasErrors() {
		errs := monsterIdErr.Concat(hunterIdErr)
		helpers.Response(c, nil, &errs)
	} else {
		result, err := hunter.GetMaterialFromMonsterUseCase(hunterId, monsterId)
		helpers.Response(c, result, err)
	}
}
