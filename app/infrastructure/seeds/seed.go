package seeds

import "yu-croco/ddd_on_golang/app/infrastructure/dao"

var materials = []dao.MonsterMaterial{
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

var materials2 = []dao.MonsterMaterial{
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

var MonsterSeed = dao.Monster{
	ID:           1,
	Name:         "ランポス",
	Life:         150,
	DefencePower: 100,
	OffensePower: 110,
	Materials:    materials,
}

var MonsterSeed2 = dao.Monster{
	ID:           2,
	Name:         "フルフル",
	Life:         2300,
	DefencePower: 300,
	OffensePower: 250,
	Materials:    materials2,
}

var HunterSeed = dao.Hunter{
	ID:              1,
	Name:            "新米ハンター",
	Life:            150,
	DefencePower:    100,
	OffensePower:    110,
	HuntedMaterials: []dao.MonsterMaterial{},
}

var HunterSeed2 = dao.Hunter{
	ID:              2,
	Name:            "中級ハンター",
	Life:            450,
	DefencePower:    280,
	OffensePower:    310,
	HuntedMaterials: materials2,
}
