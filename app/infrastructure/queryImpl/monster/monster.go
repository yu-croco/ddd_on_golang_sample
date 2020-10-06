package monster

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/dto"
	"yu-croco/ddd_on_golang/app/query"
)

type monsterQueryImpl struct{}

func NewMonsterQueryImpl() query.MonsterQuery {
	return &monsterQueryImpl{}
}

func (repo monsterQueryImpl) FindAll() *model.Monsters {
	db := infrastructure.GetDB()
	monsterDaos := dto.Monsters{}

	db.Preload("Materials").Find(&monsterDaos)

	return monsterDaos.ConvertToModel()
}
