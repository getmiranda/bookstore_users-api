package services

import (
	"github.com/getmiranda/bookstore_users-api/domain/users"
	"github.com/getmiranda/bookstore_users-api/utils/errors"
)

func GetUser(userId uint64) (*users.User, errors.APIError) {
	result := &users.User{ID: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user *users.User) (*users.User, errors.APIError) {
	if err := user.Validate(); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return user, nil
}
