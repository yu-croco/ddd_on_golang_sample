package repositoryImpl

import (
	"github.com/jinzhu/gorm"
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

func FindMonsterBy(db *gorm.DB, id int) *model.Monster {
	monsterDao := dao.Monster{}
	if db.First(&monsterDao, dao.Monster{ID: uint(id)}).RecordNotFound() {
	}

	return monsterDao.ConvertToModel()
}

func UpdateMonster(db *gorm.DB, monster *model.Monster) *model.Monster {
	monsterDao := dao.Monster{}

	if db.First(&monsterDao, dao.Monster{ID: uint(monster.Id)}).RecordNotFound() {
	}
	monsterDao.Life = monster.Life

	db.Save(&monsterDao)
	return monsterDao.ConvertToModel()
}
