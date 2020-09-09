package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

var (
	db *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=ddd_on_golang password=ddd_on_golang sslmode=disable")
	if err != nil {
		fmt.Println("db init error: ", err)
	}

	autoMigrate()
	execSeeds()

	fmt.Println("DB setup done!")
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		fmt.Println("db close error: ", err)
	}
}

func autoMigrate() {
	db.AutoMigrate(&dao.Hunter{})
	db.AutoMigrate(&dao.HuntedMonsterMaterial{})
}

func execSeeds(){
	db.Create(&dao.Hunter{
		ID: 1,
		Name: "新米ハンター",
		Life: 150,
		DefencePower: 100,
		OffensePower: 110,
		HuntedMaterials: []dao.HuntedMonsterMaterial{{
			ID: 1,
			Name: "ランポスの革",
			Rarity: 1,
		},
			{
				ID: 2,
				Name: "ランポスの爪",
				Rarity: 1,
			}},
	})
}