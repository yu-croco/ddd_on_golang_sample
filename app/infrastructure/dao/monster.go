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

type Monsters []Monster

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

func (monsters Monsters) ConvertToModel() *[]model.Monster {
	monsterModels := make([]model.Monster, len(monsters))

	for idx, monster := range monsters {
		model := model.Monster{
			Id:           int(monster.ID),
			Name:         monster.Name,
			Life:         monster.Life,
			DefencePower: monster.DefencePower,
			OffensePower: monster.OffensePower,
			Materials:    []model.MonsterMaterial{},
			AttackDamage: 0,
		}
		monsterModels[idx] = model
	}
	return &monsterModels
}
