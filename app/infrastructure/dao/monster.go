package dao
type Monster struct {
	ID uint   `json:"id" binding:"required"`
	Name string
	Life int
	DefencePower int
	OffensePower int
	Materials    []MonsterMaterial
}