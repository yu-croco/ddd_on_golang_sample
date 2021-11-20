package query

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
)

type MonsterQuery interface {
	FindAll() *model2.Monsters
}
