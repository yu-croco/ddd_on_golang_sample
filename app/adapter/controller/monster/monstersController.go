package monster

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure/queryImpl/monster"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
	usecase "yu-croco/ddd_on_golang/app/usecase/monster"
)

type Controller struct{}

func (ctrl Controller) Show(c *gin.Context) {
	monsterId, hunterIdErr := model.NewMonsterId(helpers.ConvertToInt(c.Param("id")))
	if hunterIdErr.HasErrors() {
		helpers.Response(c, nil, hunterIdErr)
	} else {
		repo := repositoryImpl.NewMonsterRepositoryImpl()

		dbResult, err := repo.FindById(monsterId)
		helpers.Response(c, dbResult, err)
	}
}

func (ctrl Controller) Index(c *gin.Context) {
	result := monster.NewMonsterQueryImpl().FindAll()
	helpers.Response(c, result, nil)
}

// ToDo: メソッド名とActionを統一させる
func (ctrl Controller) Attack(c *gin.Context) {
	monsterId, monsterIdErr := model.NewMonsterId(helpers.ConvertToInt(c.Param("id")))

	hunterId, hunterIdErr := model.NewHunterId(helpers.ConvertToInt(c.PostForm("hunterId")))
	if monsterIdErr.HasErrors() || hunterIdErr.HasErrors() {
		errs := monsterIdErr.Concat(*hunterIdErr)
		helpers.Response(c, nil, &errs)
	} else {
		result, err := usecase.AttackHunterUseCase(monsterId, hunterId)

		helpers.Response(c, result, err)
	}
}
