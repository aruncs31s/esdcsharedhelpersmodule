package helper

import (
	"log"
	"strconv"

	"staff/utils"

	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)


func GetJSONDataFromRequest[T any](c *gin.Context, responseHelper responsehelper.ResponseHelper) (T, bool) {
	var data T
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Printf("Error binding JSON: %v", err)
		responseHelper.BadRequest(c, utils.ErrBadRequest.Error(), utils.FixInvalidRequestData)
		return data, true
	}
	return data, false
}

// NOTE: Currently moving to requestHelper struct method , so remove this in the future.
// GetLimitAndOffset extracts pagination parameters from the request context.
// Matched to client side request.
//
// Params: c *gin.Context - The Gin context containing the request.
//
// Returns: (int, int) - The limit and offset values for pagination.
func GetLimitAndOffset(c *gin.Context) (int, int) {
	limit, _ := strconv.Atoi(c.DefaultQuery("page-size", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page-no", "1"))
	// calculate offset form page and limit
	offset := (page - 1) * limit
	return limit, offset
}
