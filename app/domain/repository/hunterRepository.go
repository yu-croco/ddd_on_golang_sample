package repository

import (
	"yu-croco/ddd_on_golang/app/domain/model"
)

type HunterRepository interface {
	FindById(id int) (*model.Hunter, error)
	Update(hunter *model.Hunter) (*model.Hunter, error)
}
