package framework

import "errors"

var _ error = &Error{}

type Error struct {
	ErrorCode int            `json:"error_code"`
	Message   string         `json:"message"`
	Fields    map[string]any `json:"fields,omitempty"`
}

func (e *Error) Error() string {
	fields := ""
	if e.Fields != nil {
		for k, v := range e.Fields {
			fields += k + ": " + v.(string) + ", "
		}
	}

	return e.Message + " (" + fields + ")"
}

func NewInternalError(message string) *Error {
	return &Error{
		ErrorCode: 500,
		Message:   message,
	}
}

func NewValidationError(message string) *Error {
	return &Error{
		ErrorCode: 400,
		Message:   message,
	}
}

func NewNotFoundError(message string) *Error {
	return &Error{
		ErrorCode: 404,
		Message:   message,
	}
}

func NewUnauthorizedError(message string) *Error {
	return &Error{
		ErrorCode: 401,
		Message:   message,
	}
}

func ExtErrorCode(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.ErrorCode
	}

	return 500
}

func ExtErrorMessage(err error) string {
	var e *Error
	if errors.As(err, &e) {
		return e.Message
	}

	return "Internal server error"
}

func ExtErrorFields(err error) map[string]any {
	var e *Error
	if errors.As(err, &e) {
		return e.Fields
	}

	return nil
}

func (e *Error) WithFields(fields map[string]any) *Error {
	e.Fields = fields
	return e
}

type Fields map[string]any
