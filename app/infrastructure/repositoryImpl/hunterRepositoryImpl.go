package repositoryImpl

import (
	"yu-croco/ddd_on_golang/app/domain/model/hunter"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

func FindById(id hunter.Id) *hunter.Hunter {
	var hunterDao *dao.Hunter
	infrastructure.GetDB().First(&hunterDao, int(id))

	return hunterDao.ConvertToModel()
}

func Update(hunter dao.Hunter) *hunter.Hunter {
	var hunterDao *dao.Hunter
	db := infrastructure.GetDB()
	db.First(&hunterDao, int(hunter.ID))
	hunterDao.Life = hunter.Life

	db.Save(&hunterDao)
	return hunterDao.ConvertToModel()
}
