package hunter

import (
	"github.com/jinzhu/gorm"
	"yu-croco/ddd_on_golang/app/domain/model/hunter"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

func FindById(db *gorm.DB, id hunter.Id) *hunter.Hunter {
	var hunterDao *dao.Hunter
	db.First(&hunterDao, int(id))

	return hunterDao.ConvertToModel()
}

func Update(db *gorm.DB, hunter dao.Hunter) *hunter.Hunter {
	var hunterDao *dao.Hunter

	db.First(&hunterDao, int(hunter.ID))
	hunterDao.Life = hunter.Life

	db.Save(&hunterDao)
	return hunterDao.ConvertToModel()
}
