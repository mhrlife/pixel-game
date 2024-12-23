package apperror

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAppError_WrapItsOwnKind(t *testing.T) {
	err1 := NewInternalError("1")
	err2 := NewValidationError("2").Wrap(err1)

	require.Equal(t, "2", ExtErrorMessage(err2))
	require.Equal(t, "1", ExtErrorMessage(ExtErrorWrapped(err2)))
}
