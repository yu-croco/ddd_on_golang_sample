package repositoryImpl

import (
	"github.com/jinzhu/gorm"
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

// ToDo: 循環参照をどうにかする
func FindHunterById(db *gorm.DB, id int) *model.Hunter {
	hunterDao := dao.Hunter{}
	if db.Find(&hunterDao, dao.Hunter{ID: uint(id)}).RecordNotFound() {
	}

	return hunterDao.ConvertToModel()
}

// ToDo: 循環参照をどうにかする
func UpdateHunter(db *gorm.DB, hunter dao.Hunter) *model.Hunter {
	var hunterDao *dao.Hunter

	if db.First(&hunterDao, int(hunter.ID)).RecordNotFound() {
	}

	hunterDao.Life = hunter.Life
	db.Save(&hunterDao)
	return hunterDao.ConvertToModel()
}
