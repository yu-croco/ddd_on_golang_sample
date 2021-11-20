package repository

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
	errors2 "yu-croco/ddd_on_golang/pkg/errors"
)

type HunterRepository interface {
	FindById(id model2.HunterId) (*model2.Hunter, *errors2.AppError)
	Update(hunter *model2.Hunter) (*model2.Hunter, *errors2.AppError)
	AddMonsterMaterial(hunter *model2.Hunter, material *model2.MonsterMaterial) *errors2.AppError
}
