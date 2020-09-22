package model

type HuntedMonsterMaterial struct {
	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
}

type HuntedMonsterMaterials []HuntedMonsterMaterial
