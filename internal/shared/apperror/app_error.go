package shared

import "net/http"

type AppErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *AppErr) Error() string {
	return e.Message
}

func NewAppErr(message, err string, code int) *AppErr {
	return &AppErr{
		Message: message,
		Err:     err,
		Code:    code,
	}
}

func NewBadRequestError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewConflictError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Err:     "conflict",
		Code:    http.StatusConflict,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *AppErr {
	return &AppErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewInternalServerError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewUnauthorizedError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}