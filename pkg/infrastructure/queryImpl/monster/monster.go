package monster

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
	infrastructure2 "yu-croco/ddd_on_golang/pkg/infrastructure"
	dto2 "yu-croco/ddd_on_golang/pkg/infrastructure/dto"
	query2 "yu-croco/ddd_on_golang/pkg/query"
)

type monsterQueryImpl struct{}

func NewMonsterQueryImpl() query2.MonsterQuery {
	return &monsterQueryImpl{}
}

func (repo monsterQueryImpl) FindAll() *model2.Monsters {
	db := infrastructure2.GetDB()
	monsterDaos := dto2.Monsters{}

	db.Preload("Materials").Find(&monsterDaos)

	return monsterDaos.ConvertToModel()
}
