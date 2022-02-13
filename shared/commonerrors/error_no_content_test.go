package commonerrors_test

import (
	"fmt"
	"testing"

	"github.com/nelsonlpco/transactions/shared/commonerrors"
	"github.com/stretchr/testify/require"
)

func Test_ErrorNoContentHandler(t *testing.T) {
	errorMessage := "resource not found"
	noContentError := commonerrors.NewErrorNoContent(errorMessage)

	t.Run("should be create a no content error", func(t *testing.T) {
		require.NotNil(t, noContentError)
	})

	t.Run("should be create an valid error message", func(t *testing.T) {
		expectedErrorMessage := fmt.Sprintf(`"%s"`, errorMessage)
		require.Equal(t, expectedErrorMessage, noContentError.Error())
	})
}
