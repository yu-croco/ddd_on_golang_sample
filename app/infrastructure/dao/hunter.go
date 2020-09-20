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
