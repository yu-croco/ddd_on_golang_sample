package repository

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/errors"
)

type MonsterRepository interface {
	FindById(id int) (*model.Monster, *errors.AppError)
	Update(monster *model.Monster) (*model.Monster, *errors.AppError)
}
