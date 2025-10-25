package helper

import (
	"strconv"

	sharedHelper "github.com/aruncs31s/esdcsharedhelpersmodule/interface/helper"
	"github.com/aruncs31s/esdcsharedhelpersmodule/utils"

	"github.com/gin-gonic/gin"
)

type requestHelper struct{}

func NewRequestHelper() sharedHelper.RequestHelper {
	return &requestHelper{}
}

func (r *requestHelper) GetAndValidateUsername(c *gin.Context, h sharedHelper.HasBothValidatorAndResponseHelper) (string, bool) {
	username := c.GetString("username")
	failed := validateUsername(h, username, c)
	if failed {
		return "", true
	}
	return username, false
}

// GetLimitAndOffset extracts pagination parameters from the request context.
// Matched to client side request.
//
// Params: c *gin.Context - The Gin context containing the request.
//
// Returns: (int, int) - The limit and offset values for pagination.
func (r *requestHelper) GetLimitAndOffset(c *gin.Context) (int, int) {
	limit, _ := strconv.Atoi(c.DefaultQuery("per-page", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	// calculate offset form page and limit
	offset := (page - 1) * limit
	return limit, offset
}

func validateUsername[T sharedHelper.HasBothValidatorAndResponseHelper](h T, username string, c *gin.Context) bool {
	if err := h.GetValidator().ValidateUsername(username); err != nil {
		h.GetResponseHelper().BadRequest(c, err.Error(), utils.FixInvalidUsername)
		return true
	}
	return false
}

func (r *requestHelper) GetURLParam(c *gin.Context, paramName string) string {
	return c.Param(paramName)
}

// Generic ID validation helper
func (r *requestHelper) ValidateAndParseID(h sharedHelper.HasBothValidatorAndResponseHelper, idName string, c *gin.Context, fixMessage string) (uint, bool) {
	validator := h.GetValidator()
	idStr := c.Param(idName)
	id, err := validator.ValidateIDAndParse(idStr)
	if err != nil {
		h.GetResponseHelper().BadRequest(c, err.Error(), fixMessage)
		return 0, true
	}
	return id, false
}
