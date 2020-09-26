package dto

type MonsterMaterial struct {
	ID        uint
	Name      string
	Rarity    int
	Monster   Monster `json:"-" binding:"required"`
	MonsterID uint    `gorm:"not null"  json:"monster_id"`
	Hunters   Hunters `gorm:"many2many:hunted_monster_materials"`
}

type MonsterMaterials []MonsterMaterial
