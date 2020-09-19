package dao

import "yu-croco/ddd_on_golang/app/domain/model/hunter"

type Hunter struct {
	ID              uint `json:"id" binding:"required"`
	Name            string
	Life            int
	DefencePower    int
	OffensePower    int
	HuntedMaterials []HuntedMonsterMaterial
}

func (h *Hunter) ConvertToModel() *hunter.Hunter {
	return &hunter.Hunter{
		Id:              hunter.Id(int(h.ID)),
		Name:            hunter.Name(h.Name),
		Life:            hunter.Life(h.Life),
		DefencePower:    hunter.DefencePower(h.DefencePower),
		OffensePower:    hunter.OffensePower(h.OffensePower),
		HuntedMaterials: []hunter.HuntedMonsterMaterial{},
		AttackDamage:    hunter.AttackDamage(0),
	}
}
