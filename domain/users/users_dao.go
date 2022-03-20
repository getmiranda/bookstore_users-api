package users

import (
	"fmt"
	"strings"

	bookstore_users_db "github.com/getmiranda/bookstore_users-api/datasources/mysql/bookstore_users"
	"github.com/getmiranda/bookstore_users-api/logger"
	"github.com/getmiranda/bookstore_users-api/utils/errors"
	"github.com/getmiranda/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser             = "INSERT INTO users (first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser             = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?"
)

func (user *User) Get() error {
	stmt, err := bookstore_users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.ErrorWithError("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("error when tying to get user: database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to get user by id")
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Save() error {
	stmt, err := bookstore_users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.ErrorWithError("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		logger.Error("error when trying to save user")
		return mysql_utils.ParseError(err)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.ErrorWithError("error when trying to get last insert id after creating user", err)
		return errors.NewInternalServerError("database error")
	}

	user.ID = uint64(userId)
	return nil
}

func (user *User) Update() error {
	stmt, err := bookstore_users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.ErrorWithError("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() error {
	stmt, err := bookstore_users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.ErrorWithError("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) (Users, error) {
	stmt, err := bookstore_users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.ErrorWithError("error when trying to prepare get users statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find users by status")
		return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row into user struct")
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status '%s'", status))
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() error {
	stmt, err := bookstore_users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement")
		return errors.NewInternalServerError("error when tying to find user")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by email and password")
		return errors.NewInternalServerError("error when tying to find user")
	}
	return nil
}
