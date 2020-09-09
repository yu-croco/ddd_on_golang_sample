package dao

type HuntedMonsterMaterial struct {
	ID                uint            `json:"id" binding:"required"`
	MonsterMaterial   MonsterMaterial `json:"-" binding:"required"`
	MonsterMaterialID uint            `gorm:"not null"  json:"monster_material_id"`
}
