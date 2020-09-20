package query

import "yu-croco/ddd_on_golang/app/domain/model"

type HunterQuery interface {
	FindAll() *[]model.Hunter
}
