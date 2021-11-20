package dto

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
)

type Monster struct {
	ID           uint `json:"id" binding:"required"`
	Name         string
	Life         int
	DefencePower int
	OffensePower int
	Materials    MonsterMaterials
}

type Monsters []Monster

func (m *Monster) ConvertToModel() *model2.Monster {
	var materials = make(model2.MonsterMaterials, len(m.Materials))
	for idx, material := range m.Materials {
		materials[idx] = model2.MonsterMaterial{
			Name:   model2.MonsterMaterialName(material.Name),
			Rarity: model2.MonsterMaterialRarity(material.Rarity),
		}
	}

	return &model2.Monster{
		Id:           model2.MonsterId(m.ID),
		Name:         model2.MonsterName(m.Name),
		Life:         model2.MonsterLife(m.Life),
		DefencePower: model2.MonsterDefencePower(m.DefencePower),
		OffensePower: model2.MonsterOffensePower(m.OffensePower),
		Materials:    materials,
		AttackDamage: model2.MonsterAttackDamage(0),
	}
}

func (monsters Monsters) ConvertToModel() *model2.Monsters {
	monsterModels := make(model2.Monsters, len(monsters))

	for idx, monster := range monsters {
		model := model2.Monster{
			Id:           model2.MonsterId(monster.ID),
			Name:         model2.MonsterName(monster.Name),
			Life:         model2.MonsterLife(monster.Life),
			DefencePower: model2.MonsterDefencePower(monster.DefencePower),
			OffensePower: model2.MonsterOffensePower(monster.OffensePower),
			Materials:    convertMonsterMaterialRowToModel(monster),
			AttackDamage: model2.MonsterAttackDamage(0),
		}
		monsterModels[idx] = model
	}
	return &monsterModels
}

func convertMonsterMaterialRowToModel(monster Monster) model2.MonsterMaterials {
	materials := make(model2.MonsterMaterials, len(monster.Materials))
	for idx2, material := range monster.Materials {
		materials[idx2] = model2.MonsterMaterial{
			Name:   model2.MonsterMaterialName(material.Name),
			Rarity: model2.MonsterMaterialRarity(material.Rarity),
		}
	}

	return materials
}
