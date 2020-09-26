package model

type HuntedMonsterMaterial struct {
	Name   HuntedMonsterMaterialName   `json:"name"`
	Rarity HuntedMonsterMaterialRarity `json:"rarity"`
}
type HuntedMonsterMaterialName string
type HuntedMonsterMaterialRarity int

type HuntedMonsterMaterials []HuntedMonsterMaterial
