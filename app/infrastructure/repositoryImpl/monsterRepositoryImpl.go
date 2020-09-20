package repositoryImpl

import (
	"github.com/jinzhu/gorm"
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

type MonsterRepositoryImpl struct{}

func (repositoryImpl *MonsterRepositoryImpl) FindById(db *gorm.DB, id int) *model.Monster {
	monsterDao := dao.Monster{}
	if db.First(&monsterDao, dao.Monster{ID: uint(id)}).RecordNotFound() {
	}

	return monsterDao.ConvertToModel()
}

func (repositoryImpl *MonsterRepositoryImpl) Update(db *gorm.DB, monster *model.Monster) *model.Monster {
	monsterDao := dao.Monster{}

	if db.First(&monsterDao, dao.Monster{ID: uint(monster.Id)}).RecordNotFound() {
	}
	monsterDao.Life = monster.Life

	db.Save(&monsterDao)
	return monsterDao.ConvertToModel()
}
