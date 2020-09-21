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
	hunterId := helpers.ConvertToInt(c.Param("id"))

	repo := repositoryImpl.NewHunterRepositoryImpl()
	result, errs := repo.FindById(hunterId)

	helpers.Response(c, result, errs)
}

// ToDo: メソッド名とActionを統一させる
func (ctrl Controller) Attack(c *gin.Context) {
	var monster model.Monster
	c.BindJSON(&monster)

	hunterId := helpers.ConvertToInt(c.Param("id"))
	result, errs := hunter.AttackMonsterUseCase(hunterId, monster.Id)

	helpers.Response(c, result, errs)
}

func (ctrl Controller) Index(c *gin.Context) {
	result := queryImpl.NewHunterQueryImpl().FindAll()
	helpers.Response(c, result, nil)
}

// ToDo: メソッド名とActionを統一させる
func (ctrl Controller) GetMaterial(c *gin.Context) {
	var monster model.Monster
	c.BindJSON(&monster)
	hunterId := helpers.ConvertToInt(c.Param("id"))

	result, errs := hunter.GetMaterialFromMonsterUseCase(hunterId, monster.Id)

	helpers.Response(c, result, errs)
}
