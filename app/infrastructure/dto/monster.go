package dto

import "yu-croco/ddd_on_golang/app/domain/model"

type Monster struct {
	ID           uint `json:"id" binding:"required"`
	Name         string
	Life         int
	DefencePower int
	OffensePower int
	Materials    MonsterMaterials
}

type Monsters []Monster

func (m *Monster) ConvertToModel() *model.Monster {
	var materials = make(model.MonsterMaterials, len(m.Materials))
	for idx, material := range m.Materials {
		materials[idx] = model.MonsterMaterial{
			Name:   model.MonsterMaterialName(material.Name),
			Rarity: model.MonsterMaterialRarity(material.Rarity),
		}
	}

	return &model.Monster{
		Id:           model.MonsterId(m.ID),
		Name:         model.MonsterName(m.Name),
		Life:         model.MonsterLife(m.Life),
		DefencePower: model.MonsterDefencePower(m.DefencePower),
		OffensePower: model.MonsterOffensePower(m.OffensePower),
		Materials:    materials,
		AttackDamage: model.MonsterAttackDamage(0),
	}
}

func (monsters Monsters) ConvertToModel() *model.Monsters {
	monsterModels := make(model.Monsters, len(monsters))

	for idx, monster := range monsters {
		model := model.Monster{
			Id:           model.MonsterId(monster.ID),
			Name:         model.MonsterName(monster.Name),
			Life:         model.MonsterLife(monster.Life),
			DefencePower: model.MonsterDefencePower(monster.DefencePower),
			OffensePower: model.MonsterOffensePower(monster.OffensePower),
			Materials:    convertMonsterMaterialRowToModel(monster),
			AttackDamage: model.MonsterAttackDamage(0),
		}
		monsterModels[idx] = model
	}
	return &monsterModels
}

func convertMonsterMaterialRowToModel(monster Monster) model.MonsterMaterials {
	materials := make(model.MonsterMaterials, len(monster.Materials))
	for idx2, material := range monster.Materials {
		materials[idx2] = model.MonsterMaterial{
			Name:   model.MonsterMaterialName(material.Name),
			Rarity: model.MonsterMaterialRarity(material.Rarity),
		}
	}

	return materials
}
