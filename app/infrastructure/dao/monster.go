package dao

import (
	"yu-croco/ddd_on_golang/app/domain/model/monster"
)

type Monster struct {
	ID           uint `json:"id" binding:"required"`
	Name         string
	Life         int
	DefencePower int
	OffensePower int
	Materials    []MonsterMaterial
}

func (m *Monster) ConvertToModel() *monster.Monster {
	return &monster.Monster{
		Id:           monster.Id(int(m.ID)),
		Name:         monster.Name(m.Name),
		Life:         monster.Life(m.Life),
		DefencePower: monster.DefencePower(m.DefencePower),
		OffensePower: monster.OffensePower(m.OffensePower),
		Materials:    []monster.MonsterMaterial{},
		AttackDamage: monster.AttackDamage(0),
	}
}
