package repositoryImpl

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/domain/repository"
	"yu-croco/ddd_on_golang/app/errors"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/dto"
)

type hunterRepositoryImpl struct{}

func NewHunterRepositoryImpl() repository.HunterRepository {
	return &hunterRepositoryImpl{}
}

func (repositoryImpl *hunterRepositoryImpl) FindById(id model.HunterId) (*model.Hunter, *errors.AppError) {
	db := infrastructure.GetDB()
	var err errors.AppError
	hunterDao := dto.Hunter{}

	if db.Find(&hunterDao, dto.Hunter{ID: uint(id)}).RecordNotFound() {
		err = notFoundHunterError(id)
		return nil, &err
	}

	return hunterDao.ConvertToModel(), nil
}

func (repositoryImpl *hunterRepositoryImpl) Update(hunter *model.Hunter) (*model.Hunter, *errors.AppError) {
	db := infrastructure.GetDB()
	var err errors.AppError
	hunterDao := dto.Hunter{}

	if db.First(&hunterDao, dto.Hunter{ID: uint(hunter.Id)}).RecordNotFound() {
		err = notFoundHunterError(hunter.Id)
		return nil, &err
	}

	hunterDao.Life = int(hunter.Life)
	db.Save(&hunterDao)

	return hunterDao.ConvertToModel(), nil
}

func (repositoryImpl *hunterRepositoryImpl) AddMonsterMaterial(hunter *model.Hunter, material *model.MonsterMaterial) *errors.AppError {
	db := infrastructure.GetDB()
	var err errors.AppError
	hunterDao := dto.Hunter{}

	if db.First(&hunterDao, dto.Hunter{ID: uint(hunter.Id)}).RecordNotFound() {
		err = notFoundHunterError(hunter.Id)
		return &err
	}

	var materialDao dto.MonsterMaterial
	if db.First(&materialDao, dto.MonsterMaterial{Name: string(material.Name)}).RecordNotFound() {
		err = notFoundMaterialError(string(material.Name))
		return &err
	}

	// ToDo: 同じハンターが同じ素材を複数取得できないので修正したい
	db.Model(&hunterDao).Association("HuntedMaterials").Append(materialDao)

	return nil
}

func notFoundHunterError(id model.HunterId) errors.AppError {
	return errors.NewAppError("id " + string(id) + "のhunterは見つかりませんでした")
}

func notFoundMaterialError(material string) errors.AppError {
	return errors.NewAppError(material + "見つかりませんでした")
}
