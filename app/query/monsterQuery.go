package query

import (
	"yu-croco/ddd_on_golang/app/domain/model"
)

type MonsterQuery interface {
	FindAll() *[]model.Monster
}
