package helpers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"yu-croco/ddd_on_golang/app/errors"
)

func ConvertToInt(modelId string) int {
	id, _ := strconv.Atoi(modelId)
	return id
}

func Response(c *gin.Context, result interface{}, err *errors.AppError) {
	if err.HasErrors() {
		c.JSON(400, err)
	} else {
		c.JSON(200, result)
	}
}
