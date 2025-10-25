package helper

import (
	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

type RequestHelper interface {
	// GetUsername extracts the username from the request context.
	//
	// Parameters:
	//   - c: Gin context containing the request
	//   - h: Interface providing access to validator and response helper
	//
	// Returns:
	//   - string: Extracted username
	//   - bool: True if extraction failed, false if successful
	GetAndValidateUsername(c *gin.Context, h HasBothValidatorAndResponseHelper) (string, bool)
	// ValidateAndParseID validates and parses a string ID into a uint
	// Parameters:
	//   - h: Interface providing access to validator and response helper
	//   - idStr: String ID to validate and parse
	//   - c: Gin context for handling the response
	//   - fixMessage: Message to display if validation fails
	// Returns:
	//   - uint: Parsed ID value
	//   - bool: True if validation failed, false if successful
	ValidateAndParseID(h HasBothValidatorAndResponseHelper, idName string, c *gin.Context, fixMessage string) (uint, bool)
	// GetLimitAndOffset extracts pagination parameters from the request context.
	// Matched to client side request.
	//
	// Params: c *gin.Context - The Gin context containing the request.
	//
	// Returns: (int, int) - The limit and offset values for pagination.
	GetLimitAndOffset(c *gin.Context) (int, int)
}

type HasValidator interface {
	GetValidator() RequestValidator
}

type HasResponseHelper interface {
	GetResponseHelper() responsehelper.ResponseHelper
}
type HasBothValidatorAndResponseHelper interface {
	HasValidator
	HasResponseHelper
}
