package repository

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/errors"
)

type HunterRepository interface {
	FindById(id int) (*model.Hunter, *errors.AppError)
	Update(hunter *model.Hunter) (*model.Hunter, *errors.AppError)
}
