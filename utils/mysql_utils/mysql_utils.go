package mysql_utils

import (
	"strings"

	"github.com/getmiranda/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) errors.APIError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record found", err.Error())
		}
		return errors.NewInternalServerError("error parsing database response", err.Error())
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data", sqlErr.Message)
	}
	return errors.NewInternalServerError("error processing request", sqlErr.Message)
}
