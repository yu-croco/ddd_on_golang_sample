package hunter

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

// ToDo: 循環参照をどうにかする
func FindById(db *gorm.DB, id int) {
	hunterDao := dao.Hunter{}
	if db.Find(&hunterDao, dao.Hunter{ID: uint(id)}).RecordNotFound() {
		fmt.Println("not foundddddddddddddddddddddd")
	}

	fmt.Println(hunterDao)

	//return hunterDao.ConvertToModel()
}

// ToDo: 循環参照をどうにかする
func Update(db *gorm.DB, hunter dao.Hunter) {
	var hunterDao *dao.Hunter

	if db.First(&hunterDao, int(hunter.ID)).RecordNotFound() {
	}

	hunterDao.Life = hunter.Life
	db.Save(&hunterDao)
	//return hunterDao.ConvertToModel()
}
