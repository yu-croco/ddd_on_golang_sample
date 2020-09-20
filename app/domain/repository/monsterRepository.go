package repository

import (
	"yu-croco/ddd_on_golang/app/domain/model"
)

type MonsterRepository interface {
	FindById(id int) (*model.Monster, error)
	Update(monster *model.Monster) (*model.Monster, error)
}
