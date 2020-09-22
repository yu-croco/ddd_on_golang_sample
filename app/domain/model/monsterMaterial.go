package model

type MonsterMaterial struct {
	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
}

type MonsterMaterials []MonsterMaterial
