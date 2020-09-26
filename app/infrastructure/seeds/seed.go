package seeds

import "yu-croco/ddd_on_golang/app/infrastructure/dto"

var materials = dto.MonsterMaterials{
	{
		ID:     1,
		Name:   "ランポスの皮",
		Rarity: 1,
	},
	{
		ID:     2,
		Name:   "ランポスの爪",
		Rarity: 1,
	},
}

var materials2 = dto.MonsterMaterials{
	{
		ID:     3,
		Name:   "アルビノの皮",
		Rarity: 2,
	},
	{
		ID:     4,
		Name:   "電気袋",
		Rarity: 2,
	},
}

var MonsterSeed = dto.Monster{
	ID:           1,
	Name:         "ランポス",
	Life:         150,
	DefencePower: 100,
	OffensePower: 110,
	Materials:    materials,
}

var MonsterSeed2 = dto.Monster{
	ID:           2,
	Name:         "フルフル",
	Life:         2300,
	DefencePower: 300,
	OffensePower: 250,
	Materials:    materials2,
}

var HunterSeed = dto.Hunter{
	ID:              1,
	Name:            "新米ハンター",
	Life:            150,
	DefencePower:    100,
	OffensePower:    110,
	HuntedMaterials: dto.MonsterMaterials{},
}

var HunterSeed2 = dto.Hunter{
	ID:              2,
	Name:            "中級ハンター",
	Life:            450,
	DefencePower:    280,
	OffensePower:    310,
	HuntedMaterials: materials2,
}
