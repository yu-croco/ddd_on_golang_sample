package repositoryImpl

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

type HunterRepositoryImpl struct{}

func (repositoryImpl *HunterRepositoryImpl) FindById(id int) *model.Hunter {
	db := infrastructure.GetDB()
	hunterDao := dao.Hunter{}
	if db.Find(&hunterDao, dao.Hunter{ID: uint(id)}).RecordNotFound() {
	}

	return hunterDao.ConvertToModel()
}

func (repositoryImpl *HunterRepositoryImpl) Update(hunter dao.Hunter) *model.Hunter {
	db := infrastructure.GetDB()
	var hunterDao *dao.Hunter

	if db.First(&hunterDao, int(hunter.ID)).RecordNotFound() {
	}

	hunterDao.Life = hunter.Life
	db.Save(&hunterDao)
	return hunterDao.ConvertToModel()
}
