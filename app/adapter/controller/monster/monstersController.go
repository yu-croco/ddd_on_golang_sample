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
	monsterId := helpers.ConvertToInt(c.Param("id"))

	repo := repositoryImpl.NewMonsterRepositoryImpl()
	result, errs := repo.FindById(monsterId)

	helpers.Response(c, result, errs)
}

func (ctrl Controller) Index(c *gin.Context) {
	result := monster.NewMonsterQueryImpl().FindAll()
	helpers.Response(c, result, nil)
}

// ToDo: メソッド名とActionを統一させる
func (ctrl Controller) Attack(c *gin.Context) {
	var hunter model.Hunter
	c.BindJSON(&hunter)
	monsterId := helpers.ConvertToInt(c.Param("id"))

	result, errs := usecase.AttackHunterUseCase(monsterId, hunter.Id)

	helpers.Response(c, result, errs)
}
