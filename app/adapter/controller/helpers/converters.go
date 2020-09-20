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

// Note: interface使うと型が壊れるのでアレだけど、そこはrequest testを作ることで担保する形でも良いかも
func Response(c *gin.Context, result interface{}, err *errors.AppError) {
	if err.HasErrors() {
		c.JSON(400, err)
	} else {
		c.JSON(200, result)
	}
}
