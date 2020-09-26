package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"yu-croco/ddd_on_golang/app/infrastructure/dto"
	"yu-croco/ddd_on_golang/app/infrastructure/seeds"
)

var (
	db  *gorm.DB
	err error
)

func Init() *gorm.DB {
	config := "host=db port=5432 user=postgres dbname=ddd_on_golang password=ddd_on_golang sslmode=disable"
	db, err = gorm.Open("postgres", config)
	if err != nil {
		fmt.Println("db init error: ", err)
	}

	autoMigrate()
	execSeeds()

	fmt.Println("[INFO]DB setup done!")
	return db
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
	db.
		AutoMigrate(&dto.Monster{}).
		AutoMigrate(&dto.MonsterMaterial{}).
		AutoMigrate(&dto.Hunter{}).
		AutoMigrate(&dto.HuntedMonsterMaterial{})
}

func execSeeds() {
	db.
		Create(&seeds.MonsterSeed).
		Create(&seeds.MonsterSeed2).
		Create(&seeds.HunterSeed).
		Create(&seeds.HunterSeed2)
}
