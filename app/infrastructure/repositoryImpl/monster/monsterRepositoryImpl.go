package monster

import (
	"github.com/jinzhu/gorm"
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

// ToDo: 循環参照をどうにかする
func FindBy(db *gorm.DB, id model.Id) {
	var monsterDao *dao.Monster
	if db.First(&monsterDao, int(id)).RecordNotFound() {
	}

	//return monsterDao.ConvertToModel()
}

// ToDo: 循環参照をどうにかする
func Update(db *gorm.DB, monster dao.Monster) {
	var monsterDao *dao.Monster

	if db.First(&monsterDao, int(monster.ID)).RecordNotFound() {
	}
	monsterDao.Life = monster.Life

	db.Save(&monsterDao)
	//return monsterDao.ConvertToModel()
}
