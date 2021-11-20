package dto

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
)

type Hunter struct {
	ID              uint `json:"id" binding:"required"`
	Name            string
	Life            int
	DefencePower    int
	OffensePower    int
	HuntedMaterials MonsterMaterials `gorm:"many2many:hunted_monster_materials"`
}

type Hunters []Hunter

func (h *Hunter) ConvertToModel() *model2.Hunter {
	return &model2.Hunter{
		Id:              model2.HunterId(h.ID),
		Name:            model2.HunterName(h.Name),
		Life:            model2.HunterLife(h.Life),
		DefencePower:    model2.HunterDefencePower(h.DefencePower),
		OffensePower:    model2.HunterOffensePower(h.OffensePower),
		HuntedMaterials: model2.HuntedMonsterMaterials{},
		AttackDamage:    model2.HunterAttackDamage(0),
	}
}

func (hunters Hunters) ConvertToModel() *model2.Hunters {
	result := make(model2.Hunters, len(hunters))

	for idx, hunter := range hunters {
		hunterModel := model2.Hunter{
			Id:              model2.HunterId(hunter.ID),
			Name:            model2.HunterName(hunter.Name),
			Life:            model2.HunterLife(hunter.Life),
			DefencePower:    model2.HunterDefencePower(hunter.DefencePower),
			OffensePower:    model2.HunterOffensePower(hunter.OffensePower),
			HuntedMaterials: convertMaterialRowToModel(hunter),
			AttackDamage:    model2.HunterAttackDamage(0),
		}
		result[idx] = hunterModel
	}

	return &result
}

func convertMaterialRowToModel(hunter Hunter) model2.HuntedMonsterMaterials {
	materials := make(model2.HuntedMonsterMaterials, len(hunter.HuntedMaterials))
	for idx2, material := range hunter.HuntedMaterials {
		materials[idx2] = model2.HuntedMonsterMaterial{
			Name:   model2.HuntedMonsterMaterialName(material.Name),
			Rarity: model2.HuntedMonsterMaterialRarity(material.Rarity),
		}
	}

	return materials
}
