package users

import (
	"fmt"

	bookstore_users_db "github.com/getmiranda/bookstore_users-api/datasources/mysql/bookstore_users"
	"github.com/getmiranda/bookstore_users-api/utils/date_utils"
	"github.com/getmiranda/bookstore_users-api/utils/errors"
	"github.com/getmiranda/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser       = "INSERT INTO users (first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
)

func (user *User) Get() errors.APIError {
	stmt, err := bookstore_users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError("error preparing statement", err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Save() errors.APIError {
	stmt, err := bookstore_users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("error preparing statement", err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNow()
	user.Status = StatusActive

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user %d", user.ID), err.Error())
	}

	user.ID = uint64(userId)
	return nil
}

func (user *User) Update() errors.APIError {
	stmt, err := bookstore_users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError("error preparing statement", err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() errors.APIError {
	stmt, err := bookstore_users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError("error preparing statement", err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, errors.APIError) {
	stmt, err := bookstore_users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError("error preparing statement", err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status), "")
	}
	return results, nil
}
