package mysql_utils

import (
	"strings"

	"github.com/getmiranda/bookstore_users-api/logger"
	"github.com/getmiranda/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) errors.APIError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		// No MySQL errors
		if strings.Contains(err.Error(), errorNoRows) {
			logger.ErrorWithError("record not found", err)
			return errors.NewNotFoundError("record not found")
		}
		logger.ErrorWithError("unknown error", err)
		return errors.NewInternalServerError("database error")
	}
	// MySQL errors
	switch sqlErr.Number {
	case 1062:
		logger.ErrorWithError("duplicate record", err)
		return errors.NewBadRequestError("duplicate record")
	}
	logger.ErrorWithError("unknown error", err)
	return errors.NewInternalServerError("database error")
}
