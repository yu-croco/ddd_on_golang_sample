package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	dto2 "yu-croco/ddd_on_golang/pkg/infrastructure/dto"
	seeds2 "yu-croco/ddd_on_golang/pkg/infrastructure/seeds"
)

var (
	db  *gorm.DB
	err error
)

func Init() *gorm.DB {
	fmt.Println("waiting to db start up....")
	time.Sleep(time.Second * 5)
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

func autoMigrate() {
	db.
		AutoMigrate(&dto2.Monster{}).
		AutoMigrate(&dto2.MonsterMaterial{}).
		AutoMigrate(&dto2.Hunter{}).
		AutoMigrate(&dto2.HuntedMonsterMaterial{})
}

func execSeeds() {
	db.
		Create(&seeds2.MonsterSeed).
		Create(&seeds2.MonsterSeed2).
		Create(&seeds2.HunterSeed).
		Create(&seeds2.HunterSeed2)
}
