package repositoryImpl

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/domain/repository"
	"yu-croco/ddd_on_golang/app/errors"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

type MonsterRepositoryImpl struct{}

func NewMonsterRepositoryImpl() repository.MonsterRepository {
	return &MonsterRepositoryImpl{}
}

func (repositoryImpl *MonsterRepositoryImpl) FindById(id int) (*model.Monster, *errors.AppError) {
	db := infrastructure.GetDB()
	var err errors.AppError
	monsterDao := dao.Monster{}

	if db.First(&monsterDao, dao.Monster{ID: uint(id)}).RecordNotFound() {
		err = notFoundMonsterError(id)
		return nil, &err
	}

	return monsterDao.ConvertToModel(), nil
}

func (repositoryImpl *MonsterRepositoryImpl) Update(monster *model.Monster) (*model.Monster, *errors.AppError) {
	db := infrastructure.GetDB()
	var err errors.AppError
	monsterDao := dao.Monster{}

	if db.First(&monsterDao, dao.Monster{ID: uint(monster.Id)}).RecordNotFound() {
		err = notFoundMonsterError(monster.Id)
		return nil, &err
	}

	monsterDao.Life = monster.Life

	db.Save(&monsterDao)
	return monsterDao.ConvertToModel(), nil
}

func notFoundMonsterError(id int) errors.AppError {
	return errors.NewAppError("id " + string(id) + "のmonsterは見つかりませんでした")
}
