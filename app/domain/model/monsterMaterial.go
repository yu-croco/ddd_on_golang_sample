package model

type MonsterMaterial struct {
	Name   MonsterMaterialName   `json:name`
	Rarity MonsterMaterialRarity `json:rarity`
}

type MonsterMaterialName string
type MonsterMaterialRarity int
