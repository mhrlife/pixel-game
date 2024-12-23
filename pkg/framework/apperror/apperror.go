package apperror

import (
	"fmt"
	"runtime"
	"strings"
)

var _ error = &AppError{}

type AppError struct {
	Code       int
	Fields     map[string]any
	Message    string
	Wrapped    error
	StackTrace []string
}

func (a *AppError) Error() string {
	parts := make([]string, 0)

	if a.Message != "" {
		parts = append(parts, fmt.Sprintf("Message(%s)", a.Message))
	}

	if a.Wrapped != nil {
		parts = append(parts, fmt.Sprintf("Wrapped(%s)", ExtErrorMessage(a.Wrapped)))
	}

	if a.Code != 0 {
		parts = append(parts, fmt.Sprintf("Code(%d)", a.Code))
	}

	if len(a.Fields) > 0 {
		parts = append(parts, fmt.Sprintf("Fields(%v)", a.Fields))
	}

	if len(a.StackTrace) > 0 {
		parts = append(parts, fmt.Sprintf("StackTrace(%v)", a.StackTrace))
	}

	return fmt.Sprintf("AppError{%s}", strings.Join(parts, "\n"))
}

func (a *AppError) Wrap(err error) *AppError {
	a.Wrapped = err
	a.captureStackTrace()
	return a
}

func (a *AppError) WithField(key string, value any) *AppError {
	if a.Fields == nil {
		a.Fields = make(map[string]any)
	}

	a.Fields[key] = value

	return a
}

func (a *AppError) WithMessage(message string) *AppError {
	a.Message = message

	return a
}

func (a *AppError) WithCode(code int) *AppError {
	a.Code = code

	return a
}

func (a *AppError) captureStackTrace() {
	pc := make([]uintptr, 10)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	stackTrace := make([]string, 0, n)

	for i := 0; i < 10; i++ {
		frame, more := frames.Next()
		stackTrace = append(stackTrace, fmt.Sprintf("%s\n\t%s:%d", frame.Function, frame.File, frame.Line))

		if !more {
			break
		}
	}

	a.StackTrace = stackTrace
}
