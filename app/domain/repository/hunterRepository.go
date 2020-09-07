package repository

import "yu-croco/ddd_on_golang/app/domain/model/hunter"

type HunterRepository interface {
	FindById(id *hunter.Id)(*hunter.Hunter, error)
	Update(hunter *hunter.Hunter)(*hunter.Hunter, error)
}
