package helper

import (
	"fmt"

	"github.com/aruncs31s/esdcsharedhelpersmodule/interface/helper"
	"github.com/aruncs31s/esdcsharedhelpersmodule/utils"
)

type requestValidator struct{}

func NewRequestValidator() helper.RequestValidator {
	return &requestValidator{}
}

func (r *requestValidator) ValidateUsername(username string) error {
	if username == "" {
		return utils.ErrInvalidUsername
	}
	return nil
}
func (r *requestValidator) ValidateIDAndParse(param string) (uint, error) {
	var id uint
	_, err := fmt.Sscanf(param, "%d", &id)
	if err != nil || id == 0 {
		return 0, utils.ErrInvalidID
	}
	return id, nil
}
