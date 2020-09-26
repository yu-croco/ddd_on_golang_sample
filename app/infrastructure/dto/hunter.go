package dto

import (
	"yu-croco/ddd_on_golang/app/domain/model"
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

func (h *Hunter) ConvertToModel() *model.Hunter {
	return &model.Hunter{
		Id:              model.HunterId(h.ID),
		Name:            model.HunterName(h.Name),
		Life:            model.HunterLife(h.Life),
		DefencePower:    model.HunterDefencePower(h.DefencePower),
		OffensePower:    model.HunterOffensePower(h.OffensePower),
		HuntedMaterials: model.HuntedMonsterMaterials{},
		AttackDamage:    model.HunterAttackDamage(0),
	}
}

func (hunters Hunters) ConvertToModel() *model.Hunters {
	result := make(model.Hunters, len(hunters))

	for idx, hunter := range hunters {
		hunterModel := model.Hunter{
			Id:              model.HunterId(hunter.ID),
			Name:            model.HunterName(hunter.Name),
			Life:            model.HunterLife(hunter.Life),
			DefencePower:    model.HunterDefencePower(hunter.DefencePower),
			OffensePower:    model.HunterOffensePower(hunter.OffensePower),
			HuntedMaterials: convertMaterialRowToModel(hunter),
			AttackDamage:    model.HunterAttackDamage(0),
		}
		result[idx] = hunterModel
	}

	return &result
}

func convertMaterialRowToModel(hunter Hunter) model.HuntedMonsterMaterials {
	materials := make(model.HuntedMonsterMaterials, len(hunter.HuntedMaterials))
	for idx2, material := range hunter.HuntedMaterials {
		materials[idx2] = model.HuntedMonsterMaterial{
			Name:   model.HuntedMonsterMaterialName(material.Name),
			Rarity: model.HuntedMonsterMaterialRarity(material.Rarity),
		}
	}

	return materials
}
