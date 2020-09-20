package queryImpl

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
	"yu-croco/ddd_on_golang/app/query"
)

type HunterQueryImpl struct{}

func NewHunterQueryImpl() query.HunterQuery {
	return &HunterQueryImpl{}
}

func (repo HunterQueryImpl) FindAll() *[]model.Hunter {
	db := infrastructure.GetDB()
	hunterDaos := dao.Hunters{}

	db.Find(&hunterDaos)

	return hunterDaos.ConvertToModel()
}
