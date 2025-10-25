package helper

import (
	"fmt"

	sharedHelper "github.com/aruncs31s/esdcsharedhelpersmodule/interface/helper"
)

type errorHelper struct {
}

func (errorHelper) GetRecordDoesNotBelongErrorMessage(id any, user string) error {
	return fmt.Errorf("the record %d does not belong to  %v", id, user)
}
func NewErrorHelper() sharedHelper.ErrorHelper {
	return &errorHelper{}
}
