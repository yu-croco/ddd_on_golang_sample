package repositoryImpl

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/domain/repository"
	"yu-croco/ddd_on_golang/app/errors"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/dto"
)

type MonsterRepositoryImpl struct{}

func NewMonsterRepositoryImpl() repository.MonsterRepository {
	return &MonsterRepositoryImpl{}
}

func (repositoryImpl *MonsterRepositoryImpl) FindById(id model.MonsterId) (*model.Monster, *errors.AppError) {
	db := infrastructure.GetDB()
	var err errors.AppError
	monsterDao := dto.Monster{}

	if db.Preload("Materials").First(&monsterDao, dto.Monster{ID: uint(id)}).RecordNotFound() {
		err = notFoundMonsterError(id)
		return nil, &err
	}

	return monsterDao.ConvertToModel(), nil
}

func (repositoryImpl *MonsterRepositoryImpl) Update(monster *model.Monster) (*model.Monster, *errors.AppError) {
	db := infrastructure.GetDB()
	var err errors.AppError
	monsterDao := dto.Monster{}

	if db.First(&monsterDao, dto.Monster{ID: uint(monster.Id)}).RecordNotFound() {
		err = notFoundMonsterError(monster.Id)
		return nil, &err
	}

	monsterDao.Life = int(monster.Life)

	db.Save(&monsterDao)
	return monsterDao.ConvertToModel(), nil
}

func notFoundMonsterError(id model.MonsterId) errors.AppError {
	return errors.NewAppError("id " + string(id) + "のmonsterは見つかりませんでした")
}
