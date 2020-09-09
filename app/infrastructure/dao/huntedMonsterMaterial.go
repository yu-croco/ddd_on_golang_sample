package dao

type HuntedMonsterMaterial struct {
	ID uint   `json:"id" binding:"required"`
	Name string `json:name binding:"required"`
	Rarity int `json:rarity binding:"required"`
	Hunter    Hunter   `json:"-" binding:"required"`
	HunterID  uint   `gorm:"not null"  json:"hunter_id"`
}
