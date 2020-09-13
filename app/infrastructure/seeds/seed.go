package seeds

import "yu-croco/ddd_on_golang/app/infrastructure/dao"

var MonsterSeed = dao.Monster{
	ID:           1,
	Name:         "ランポス",
	Life:         150,
	DefencePower: 100,
	OffensePower: 110,
	Materials: []dao.MonsterMaterial{
		{
			ID:     1,
			Name:   "ランポスの革",
			Rarity: 1,
		},
		{
			ID:     2,
			Name:   "ランポスの爪",
			Rarity: 1,
		},
	},
}

var HunterSeed = dao.Hunter{
	ID:           1,
	Name:         "新米ハンター",
	Life:         150,
	DefencePower: 100,
	OffensePower: 110,
	HuntedMaterials: []dao.HuntedMonsterMaterial{{
		ID:                1,
		MonsterMaterialID: 1,
	},
		{
			ID:                2,
			MonsterMaterialID: 2,
		}},
}
