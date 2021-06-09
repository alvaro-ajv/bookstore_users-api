package mysql_utils

import (
	"errors"
	"github.com/alvaro259818/bookstore-utils-go/rest_errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("no records matching to given id")
		}
		return rest_errors.NewInternalServerError("error parsing database response", err)
	}
	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("invalid data, duplicated key")
	}
	return rest_errors.NewInternalServerError("error processing request", errors.New("database error"))
}
