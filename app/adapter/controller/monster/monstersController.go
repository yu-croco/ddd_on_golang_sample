package monster

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
)

type Controller struct{}

func (ctrl Controller) Show(c *gin.Context) {
	hunterId := c.Param("id")
	id, _ := strconv.Atoi(hunterId)
	repo := repositoryImpl.MonsterRepositoryImpl{}

	dbResult := repo.FindById(id)
	c.JSON(200, dbResult)
}
