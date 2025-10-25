package helper

type ErrorHelper interface {
	GetRecordDoesNotBelongErrorMessage(id any, user string) error
}
