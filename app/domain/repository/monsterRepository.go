package repository

import (
	"yu-croco/ddd_on_golang/app/domain/model/monster"
)

type MonsterRepository interface {
	FindById(id *monster.Id)(*monster.Monster, error)
	Update(monster *monster.Monster)(*monster.Monster, error)
}
