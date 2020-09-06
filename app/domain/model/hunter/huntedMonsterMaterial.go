package hunter

type HuntedMonsterMaterial struct {
	Name MaterialName `json:name`
	Rarity MaterialRarity `json:rarity`
}

type MaterialName string
type MaterialRarity int