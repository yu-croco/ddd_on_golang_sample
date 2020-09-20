package repositoryImpl

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

type MonsterRepositoryImpl struct{}

func (repositoryImpl *MonsterRepositoryImpl) FindById(id int) *model.Monster {
	db := infrastructure.GetDB()
	monsterDao := dao.Monster{}
	if db.First(&monsterDao, dao.Monster{ID: uint(id)}).RecordNotFound() {
	}

	return monsterDao.ConvertToModel()
}

func (repositoryImpl *MonsterRepositoryImpl) Update(monster *model.Monster) *model.Monster {
	db := infrastructure.GetDB()
	monsterDao := dao.Monster{}

	if db.First(&monsterDao, dao.Monster{ID: uint(monster.Id)}).RecordNotFound() {
	}
	monsterDao.Life = monster.Life

	db.Save(&monsterDao)
	return monsterDao.ConvertToModel()
}
