package dao

type Hunter struct {
	ID              uint `json:"id" binding:"required"`
	Name            string
	Life            int
	DefencePower    int
	OffensePower    int
	HuntedMaterials []MonsterMaterial `gorm:"many2many:hunted_monster_materials"`
}

//func (h *Hunter) ConvertToModel() *hunter.Hunter {
//	return &hunter.Hunter{
//		Id:              hunter.Id(int(h.ID)),
//		Name:            hunter.Name(h.Name),
//		Life:            hunter.Life(h.Life),
//		DefencePower:    hunter.DefencePower(h.DefencePower),
//		OffensePower:    hunter.OffensePower(h.OffensePower),
//		HuntedMaterials: []hunter.HuntedMonsterMaterial{},
//		AttackDamage:    hunter.AttackDamage(0),
//	}
//}
