package exception

type ServiceValidationError struct {
	Error string `json:"error"`
}

func NewServiceValidationError(error string) ServiceValidationError {
	return ServiceValidationError{
		Error: error,
	}
}
