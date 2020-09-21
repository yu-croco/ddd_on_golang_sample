package repositoryImpl

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/domain/repository"
	"yu-croco/ddd_on_golang/app/errors"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

type HunterRepositoryImpl struct{}

func NewHunterRepositoryImpl() repository.HunterRepository {
	return &HunterRepositoryImpl{}
}

func (repositoryImpl *HunterRepositoryImpl) FindById(id int) (*model.Hunter, *errors.AppError) {
	db := infrastructure.GetDB()
	var err errors.AppError
	hunterDao := dao.Hunter{}

	if db.Find(&hunterDao, dao.Hunter{ID: uint(id)}).RecordNotFound() {
		err = notFoundHunterError(id)
		return nil, &err
	}

	return hunterDao.ConvertToModel(), nil
}

func (repositoryImpl *HunterRepositoryImpl) Update(hunter *model.Hunter) (*model.Hunter, *errors.AppError) {
	db := infrastructure.GetDB()
	var err errors.AppError
	hunterDao := dao.Hunter{}

	if db.First(&hunterDao, dao.Hunter{ID: uint(hunter.Id)}).RecordNotFound() {
		err = notFoundHunterError(hunter.Id)
		return nil, &err
	}

	hunterDao.Life = hunter.Life
	db.Save(&hunterDao)

	return hunterDao.ConvertToModel(), nil
}

func (repositoryImpl *HunterRepositoryImpl) AddMonsterMaterial(hunter *model.Hunter, material *model.MonsterMaterial) *errors.AppError {
	db := infrastructure.GetDB()
	var err errors.AppError
	hunterDao := dao.Hunter{}

	if db.First(&hunterDao, dao.Hunter{ID: uint(hunter.Id)}).RecordNotFound() {
		err = notFoundHunterError(hunter.Id)
		return &err
	}

	var materialDao dao.MonsterMaterial
	if db.First(&materialDao, dao.MonsterMaterial{Name: material.Name}).RecordNotFound() {
		err = notFoundMaterialError(material.Name)
		return &err
	}

	// ToDo: 同じハンターが同じ素材を複数取得できないので修正したい
	db.Model(&hunterDao).Association("HuntedMaterials").Append(materialDao)

	return nil
}

func notFoundHunterError(id int) errors.AppError {
	return errors.NewAppError("id " + string(id) + "のhunterは見つかりませんでした")
}

func notFoundMaterialError(material string) errors.AppError {
	return errors.NewAppError(material + "見つかりませんでした")
}
