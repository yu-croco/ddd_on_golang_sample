package repositoryImpl

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
	repository2 "yu-croco/ddd_on_golang/pkg/domain/repository"
	errors2 "yu-croco/ddd_on_golang/pkg/errors"
	infrastructure2 "yu-croco/ddd_on_golang/pkg/infrastructure"
	dto2 "yu-croco/ddd_on_golang/pkg/infrastructure/dto"
)

type hunterRepositoryImpl struct{}

func NewHunterRepositoryImpl() repository2.HunterRepository {
	return &hunterRepositoryImpl{}
}

func (repositoryImpl *hunterRepositoryImpl) FindById(id model2.HunterId) (*model2.Hunter, *errors2.AppError) {
	db := infrastructure2.GetDB()
	var err errors2.AppError
	hunterDao := dto2.Hunter{}

	if db.Find(&hunterDao, dto2.Hunter{ID: uint(id)}).RecordNotFound() {
		err = notFoundHunterError(id)
		return nil, &err
	}

	return hunterDao.ConvertToModel(), nil
}

func (repositoryImpl *hunterRepositoryImpl) Update(hunter *model2.Hunter) (*model2.Hunter, *errors2.AppError) {
	db := infrastructure2.GetDB()
	var err errors2.AppError
	hunterDao := dto2.Hunter{}

	if db.First(&hunterDao, dto2.Hunter{ID: uint(hunter.Id)}).RecordNotFound() {
		err = notFoundHunterError(hunter.Id)
		return nil, &err
	}

	hunterDao.Life = int(hunter.Life)
	db.Save(&hunterDao)

	return hunterDao.ConvertToModel(), nil
}

func (repositoryImpl *hunterRepositoryImpl) AddMonsterMaterial(hunter *model2.Hunter, material *model2.MonsterMaterial) *errors2.AppError {
	db := infrastructure2.GetDB()
	var err errors2.AppError
	hunterDao := dto2.Hunter{}

	if db.First(&hunterDao, dto2.Hunter{ID: uint(hunter.Id)}).RecordNotFound() {
		err = notFoundHunterError(hunter.Id)
		return &err
	}

	var materialDao dto2.MonsterMaterial
	if db.First(&materialDao, dto2.MonsterMaterial{Name: string(material.Name)}).RecordNotFound() {
		err = notFoundMaterialError(string(material.Name))
		return &err
	}

	// ToDo: 同じハンターが同じ素材を複数取得できないので修正したい
	db.Model(&hunterDao).Association("HuntedMaterials").Append(materialDao)

	return nil
}

func notFoundHunterError(id model2.HunterId) errors2.AppError {
	return errors2.NewAppError("id " + string(id) + "のhunterは見つかりませんでした")
}

func notFoundMaterialError(material string) errors2.AppError {
	return errors2.NewAppError(material + "見つかりませんでした")
}
