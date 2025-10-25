package helper

type ErrorHelper interface {
	// GetRecordDoesNotBelongErrorMessage constructs an error message indicating that a record does not belong to a specific user.
	//
	// Params:
	// - id: any type representing the record ID.
	// - user: string representing the username.
	//
	// Returns:
	// - error: constructed error message.
	GetRecordDoesNotBelongErrorMessage(id any, user string) error
}
