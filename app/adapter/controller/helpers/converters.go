package helpers

import "strconv"

func ConvertToInt(modelId string) int {
	id, _ := strconv.Atoi(modelId)
	return id
}
