package dao

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
		Id:              int(h.ID),
		Name:            h.Name,
		Life:            h.Life,
		DefencePower:    h.DefencePower,
		OffensePower:    h.OffensePower,
		HuntedMaterials: model.HuntedMonsterMaterials{},
		AttackDamage:    0,
	}
}

func (hunters Hunters) ConvertToModel() *model.Hunters {
	result := make(model.Hunters, len(hunters))

	for idx, hunter := range hunters {
		hunterModel := model.Hunter{
			Id:              int(hunter.ID),
			Name:            hunter.Name,
			Life:            hunter.Life,
			DefencePower:    hunter.DefencePower,
			OffensePower:    hunter.OffensePower,
			HuntedMaterials: convertMaterialRowToModel(hunter),
			AttackDamage:    0,
		}
		result[idx] = hunterModel
	}

	return &result
}

func convertMaterialRowToModel(hunter Hunter) model.HuntedMonsterMaterials {
	materials := make(model.HuntedMonsterMaterials, len(hunter.HuntedMaterials))
	for idx2, material := range hunter.HuntedMaterials {
		materials[idx2] = model.HuntedMonsterMaterial{
			Name:   material.Name,
			Rarity: material.Rarity,
		}
	}

	return materials
}
