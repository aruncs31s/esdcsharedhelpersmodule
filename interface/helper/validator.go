package helper

type RequestValidator interface {
	// ValidateUsername checks if the provided username is valid by ensuring it is not empty.
	//
	// Params:
	//  - username: string representing the username to validate.
	// Returns:
	//  - error: returns an error if validation fails, nil otherwise.
	ValidateUsername(username string) error
	// ValidateID converts the provided qualification ID string to a uint after validating it.
	//
	// Params:
	//  - qualificationIDParam: string representing the qualification ID to validate and convert.
	// Returns:
	//  - uint: the converted ID.
	//  - error: returns an error if validation or conversion fails, nil otherwise.
	ValidateIDAndParse(param string) (uint, error)
}
