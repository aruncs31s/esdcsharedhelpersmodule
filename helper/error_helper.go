package helper

import (
	"esdcsharedhelpers/interface/helper"
	"fmt"
)

type errorHelper struct {
}

func (errorHelper) GetRecordDoesNotBelongErrorMessage(id any, user string) error {
	return fmt.Errorf("the record %d does not belong to  %v", id, user)
}
func NewErrorHelper() helper.ErrorHelper {
	return &errorHelper{}
}
