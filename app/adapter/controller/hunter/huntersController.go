package hunter

import (
	"github.com/gin-gonic/gin"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
)

type Controller struct{}

func (ctrl Controller) Show(c *gin.Context) {
	db := infrastructure.GetDB()
	dbResult := repositoryImpl.FindHunterById(db, 1)

	c.JSON(200, dbResult)
}
