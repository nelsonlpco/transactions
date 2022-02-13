package commonerrors_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
	"github.com/stretchr/testify/require"
)

func Test_ErrorSqlManagerHandler(t *testing.T) {
	errorMessage := "bad request"
	sqlError := commonerrors.NewErrorSql(errorMessage)

	t.Run("should be create a sql error", func(t *testing.T) {
		require.NotNil(t, sqlError)
	})

	t.Run("Should be create a valid error message", func(t *testing.T) {
		expectedMessage := fmt.Sprintf(`"%s"`, errorMessage)
		require.Equal(t, expectedMessage, sqlError.Error())
	})
}

func Test_HandleSqlErrors(t *testing.T) {
	t.Run("sholud be returns an new error when error message contains UNIQUE reference", func(t *testing.T) {
		mysqlError := errors.New("UNIQUE constraint violated")
		err := commonerrors.HandleSqlError(mysqlError)
		var expectedError *commonerrors.ErrorSql

		require.True(t, errors.As(err, &expectedError))
	})

	t.Run("sholud be returns an new error when error message contains the code 1062 reference", func(t *testing.T) {
		mysqlError := errors.New("code 1062 constraint violated")
		err := commonerrors.HandleSqlError(mysqlError)
		var expectedError *commonerrors.ErrorSql

		require.True(t, errors.As(err, &expectedError))
	})

	t.Run("should be returns an new error sql when unique is duplicated", func(t *testing.T) {
		mysqlError := &mysql.MySQLError{Number: 1169, Message: "constraint violation"}
		err := commonerrors.HandleSqlError(mysqlError)
		var expectedError *commonerrors.ErrorSql

		require.True(t, errors.As(err, &expectedError))
	})

	t.Run("should be returns an internal server error when receive an unaspected error", func(t *testing.T) {
		mysqlError := &mysql.MySQLError{Number: 1170, Message: "constraint violation"}
		err := commonerrors.HandleSqlError(mysqlError)
		var expectedError *commonerrors.ErrorInternalServer

		require.True(t, errors.As(err, &expectedError))
	})

	t.Run("should be return the same error when not a sql error", func(t *testing.T) {
		unexpectedError := errors.New("unexpected error")
		err := commonerrors.HandleSqlError(unexpectedError)

		require.Equal(t, unexpectedError, err)
	})
}
