package hunter

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
	infrastructure2 "yu-croco/ddd_on_golang/pkg/infrastructure"
	dto2 "yu-croco/ddd_on_golang/pkg/infrastructure/dto"
	query2 "yu-croco/ddd_on_golang/pkg/query"
)

type hunterQueryImpl struct{}

func NewHunterQueryImpl() query2.HunterQuery {
	return &hunterQueryImpl{}
}

func (repo hunterQueryImpl) FindAll() *model2.Hunters {
	db := infrastructure2.GetDB()
	hunterDaos := dto2.Hunters{}

	db.Preload("HuntedMaterials").Find(&hunterDaos)

	return hunterDaos.ConvertToModel()
}
