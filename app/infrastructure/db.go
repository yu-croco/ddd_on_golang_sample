package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
	"yu-croco/ddd_on_golang/app/infrastructure/seeds"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	config := "host=db port=5432 user=postgres dbname=ddd_on_golang password=ddd_on_golang sslmode=disable"
	db, err = gorm.Open("postgres", config)
	if err != nil {
		fmt.Println("db init error: ", err)
	}

	autoMigrate()
	execSeeds()

	fmt.Println("[INFO]DB setup done!")
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
	db.AutoMigrate(&dao.Monster{})
	db.AutoMigrate(&dao.MonsterMaterial{})
	db.AutoMigrate(&dao.Hunter{})
	db.AutoMigrate(&dao.HuntedMonsterMaterial{})
}

func execSeeds() {
	db.Create(&seeds.MonsterSeed).Create(&seeds.HunterSeed)
}
