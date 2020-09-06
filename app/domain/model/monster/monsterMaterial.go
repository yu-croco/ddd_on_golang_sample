package monster

type MonsterMaterial struct {
	Name   MaterialName   `json:name`
	Rarity MaterialRarity `json:rarity`
}

type MaterialName string
type MaterialRarity int