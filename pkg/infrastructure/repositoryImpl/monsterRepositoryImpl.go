package repositoryImpl

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
	repository2 "yu-croco/ddd_on_golang/pkg/domain/repository"
	errors2 "yu-croco/ddd_on_golang/pkg/errors"
	infrastructure2 "yu-croco/ddd_on_golang/pkg/infrastructure"
	dto2 "yu-croco/ddd_on_golang/pkg/infrastructure/dto"
)

type monsterRepositoryImpl struct{}

func NewMonsterRepositoryImpl() repository2.MonsterRepository {
	return &monsterRepositoryImpl{}
}

func (repositoryImpl *monsterRepositoryImpl) FindById(id model2.MonsterId) (*model2.Monster, *errors2.AppError) {
	db := infrastructure2.GetDB()
	var err errors2.AppError
	monsterDao := dto2.Monster{}

	if db.Preload("Materials").First(&monsterDao, dto2.Monster{ID: uint(id)}).RecordNotFound() {
		err = notFoundMonsterError(id)
		return nil, &err
	}

	return monsterDao.ConvertToModel(), nil
}

func (repositoryImpl *monsterRepositoryImpl) Update(monster *model2.Monster) (*model2.Monster, *errors2.AppError) {
	db := infrastructure2.GetDB()
	var err errors2.AppError
	monsterDao := dto2.Monster{}

	if db.First(&monsterDao, dto2.Monster{ID: uint(monster.Id)}).RecordNotFound() {
		err = notFoundMonsterError(monster.Id)
		return nil, &err
	}

	monsterDao.Life = int(monster.Life)

	db.Save(&monsterDao)
	return monsterDao.ConvertToModel(), nil
}

func notFoundMonsterError(id model2.MonsterId) errors2.AppError {
	return errors2.NewAppError("id " + string(id) + "のmonsterは見つかりませんでした")
}
