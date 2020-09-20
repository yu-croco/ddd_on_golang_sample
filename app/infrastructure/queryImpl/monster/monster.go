package monster

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
	"yu-croco/ddd_on_golang/app/query"
)

type MonsterQueryImpl struct{}

func NewMonsterQueryImpl() query.MonsterQuery {
	return &MonsterQueryImpl{}
}

func (repo MonsterQueryImpl) FindAll() *[]model.Monster {
	db := infrastructure.GetDB()
	monsterDaos := dao.Monsters{}

	db.Find(&monsterDaos)

	return monsterDaos.ConvertToModel()
}
