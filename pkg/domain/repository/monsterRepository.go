package repository

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
	errors2 "yu-croco/ddd_on_golang/pkg/errors"
)

type MonsterRepository interface {
	FindById(id model2.MonsterId) (*model2.Monster, *errors2.AppError)
	Update(monster *model2.Monster) (*model2.Monster, *errors2.AppError)
}
