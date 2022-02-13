package users

import (
	"fmt"

	"github.com/getmiranda/bookstore_users-api/datasources/mysql/bookstore_users"
	"github.com/getmiranda/bookstore_users-api/utils/date_utils"
	"github.com/getmiranda/bookstore_users-api/utils/errors"
	"github.com/getmiranda/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?;"
	queryUpdateUser = "UPDATE users SET first_name = ?, last_name=?, email = ? WHERE id = ?"
)

func (user *User) Get() errors.APIError {
	stmt, err := bookstore_users.ClientDB.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Save() errors.APIError {
	stmt, err := bookstore_users.ClientDB.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user %d: %s", user.ID, err.Error()))
	}

	user.ID = uint64(userId)
	return nil
}

func (user *User) Update() errors.APIError {
	stmt, err := bookstore_users.ClientDB.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
