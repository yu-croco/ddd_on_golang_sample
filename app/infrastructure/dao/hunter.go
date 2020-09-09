package dao

type Hunter struct {
	ID uint   `json:"id" binding:"required"`
	Name string
	Life int
	DefencePower int
	OffensePower int
	HuntedMaterials []HuntedMonsterMaterial
}
