package apperror

import "errors"

func ExtErrorCode(err error) int {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code
	}

	return 500
}

func ExtErrorMessage(err error) string {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Message
	}

	return err.Error()
}

func ExtErrorFields(err error) map[string]any {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Fields
	}

	return nil
}

func ExtErrorWrapped(err error) error {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Wrapped
	}

	return nil
}

func ExtErrorStackTrace(err error) []string {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.StackTrace
	}

	return nil
}

func IsConstraintError(err error) bool {
	return ExtErrorCode(err) == 409
}

func IsNotFoundError(err error) bool {
	return ExtErrorCode(err) == 404
}

func IsValidationError(err error) bool {
	return ExtErrorCode(err) == 400
}

func IsUnauthorizedError(err error) bool {
	return ExtErrorCode(err) == 401
}

func IsInternalError(err error) bool {
	return ExtErrorCode(err) == 500
}
