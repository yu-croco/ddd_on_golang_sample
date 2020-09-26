package hunter

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/dto"
	"yu-croco/ddd_on_golang/app/query"
)

type HunterQueryImpl struct{}

func NewHunterQueryImpl() query.HunterQuery {
	return &HunterQueryImpl{}
}

func (repo HunterQueryImpl) FindAll() *model.Hunters {
	db := infrastructure.GetDB()
	hunterDaos := dto.Hunters{}

	db.Preload("HuntedMaterials").Find(&hunterDaos)

	return hunterDaos.ConvertToModel()
}
