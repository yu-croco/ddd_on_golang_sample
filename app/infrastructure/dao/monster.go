package dao

import "yu-croco/ddd_on_golang/app/domain/model"

type Monster struct {
	ID           uint `json:"id" binding:"required"`
	Name         string
	Life         int
	DefencePower int
	OffensePower int
	Materials    []MonsterMaterial
}

func (m *Monster) ConvertToModel() *model.Monster {
	return &model.Monster{
		Id:           int(m.ID),
		Name:         m.Name,
		Life:         m.Life,
		DefencePower: m.DefencePower,
		OffensePower: m.OffensePower,
		Materials:    []model.MonsterMaterial{},
		AttackDamage: 0,
	}
}
