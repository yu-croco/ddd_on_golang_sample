package monster

import (
	"github.com/jinzhu/gorm"
	"yu-croco/ddd_on_golang/app/domain/model/monster"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

func FindBy(db *gorm.DB, id monster.Id) *monster.Monster {
	var monsterDao *dao.Monster
	db.First(&monsterDao, int(id))

	return monsterDao.ConvertToModel()
}

func Update(db *gorm.DB, monster dao.Monster) *monster.Monster {
	var monsterDao *dao.Monster

	db.First(&monsterDao, int(monster.ID))
	monsterDao.Life = monster.Life

	db.Save(&monsterDao)
	return monsterDao.ConvertToModel()
}
