package dao

import "yu-croco/ddd_on_golang/app/domain/model"

type Hunter struct {
	ID              uint `json:"id" binding:"required"`
	Name            string
	Life            int
	DefencePower    int
	OffensePower    int
	HuntedMaterials []MonsterMaterial `gorm:"many2many:hunted_monster_materials"`
}

type Hunters []Hunter

func (h *Hunter) ConvertToModel() *model.Hunter {
	return &model.Hunter{
		Id:              int(h.ID),
		Name:            h.Name,
		Life:            h.Life,
		DefencePower:    h.DefencePower,
		OffensePower:    h.OffensePower,
		HuntedMaterials: []model.HuntedMonsterMaterial{},
		AttackDamage:    0,
	}
}

func (hunters Hunters) ConvertToModel() *[]model.Hunter {
	result := make([]model.Hunter, len(hunters))

	for idx, hunter := range hunters {
		hunterModel := model.Hunter{
			Id:              int(hunter.ID),
			Name:            hunter.Name,
			Life:            hunter.Life,
			DefencePower:    hunter.DefencePower,
			OffensePower:    hunter.OffensePower,
			HuntedMaterials: []model.HuntedMonsterMaterial{},
			AttackDamage:    0,
		}
		result[idx] = hunterModel
	}

	return &result
}
