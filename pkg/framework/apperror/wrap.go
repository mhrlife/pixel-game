package apperror

import "github.com/mhrlife/tonference/internal/ent"

func Wrap(err error, message string) *AppError {
	switch {
	case ent.IsConstraintError(err):
		return NewConstraintError(message).Wrap(err)
	case ent.IsNotFound(err):
		return NewNotFoundError(message).Wrap(err)
	case ent.IsValidationError(err):
		return NewValidationError(message).Wrap(err)
	}

	e := &AppError{
		Wrapped: err,
		Message: message,
		Code:    500,
	}

	e.captureStackTrace()

	return e
}

func NewInternalError(msg string) *AppError {
	return &AppError{
		Code:    500,
		Message: msg,
	}
}

func NewValidationError(msg string) *AppError {
	return &AppError{
		Code:    400,
		Message: msg,
	}
}

func NewNotFoundError(msg string) *AppError {
	return &AppError{
		Code:    404,
		Message: msg,
	}
}

func NewUnauthorizedError(msg string) *AppError {
	return &AppError{
		Code:    401,
		Message: msg,
	}
}

func NewConstraintError(msg string) *AppError {
	return &AppError{
		Code:    409,
		Message: msg,
	}
}
