package exception

type ValidationError struct {
	Error string `json:"error"`
}

func NewValidationError(error string) ValidationError {
	return ValidationError{
		Error: error,
	}
}
