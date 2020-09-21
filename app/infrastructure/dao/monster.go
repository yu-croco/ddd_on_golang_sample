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
	var materials = make([]model.MonsterMaterial, len(m.Materials))
	for idx, material := range m.Materials {
		materials[idx] = model.MonsterMaterial{
			Name:   material.Name,
			Rarity: material.Rarity,
		}
	}

	return &model.Monster{
		Id:           int(m.ID),
		Name:         m.Name,
		Life:         m.Life,
		DefencePower: m.DefencePower,
		OffensePower: m.OffensePower,
		Materials:    materials,
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
			Materials:    convertMonsterMaterialRowToModel(monster),
			AttackDamage: 0,
		}
		monsterModels[idx] = model
	}
	return &monsterModels
}

func convertMonsterMaterialRowToModel(monster Monster) []model.MonsterMaterial {
	materials := make([]model.MonsterMaterial, len(monster.Materials))
	for idx2, material := range monster.Materials {
		materials[idx2] = model.MonsterMaterial{
			Name:   material.Name,
			Rarity: material.Rarity,
		}
	}

	return materials
}
