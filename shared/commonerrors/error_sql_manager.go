package commonerrors

import (
	"fmt"
	"log"
	"regexp"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
)

type ErrorSql struct {
	errorMessage string
}

func NewErrorSql(errorMessage string) *ErrorSql {
	return &ErrorSql{errorMessage: errorMessage}
}

func (e ErrorSql) Error() string {
	return fmt.Sprintf(`"%v"`, e.errorMessage)
}

func HandleSqlError(err error) error {
	if ok, _ := regexp.MatchString("UNIQUE|1062", err.Error()); ok {
		return NewErrorSql("Bad Operation")
	}

	if driverErr, ok := err.(*mysql.MySQLError); ok {
		log.Println(driverErr.Number)
		switch driverErr.Number {
		case mysqlerr.ER_DUP_UNIQUE:
			return NewErrorSql("Bad Operation")
		default:
			return NewErrorInternalServer("SQL", err.Error())
		}
	}
	return err
}
