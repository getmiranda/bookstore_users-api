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

func UpdateUser(isPartial bool, user *users.User) (*users.User, errors.APIError) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId uint64) errors.APIError {
	user := &users.User{ID: userId}
	if err := user.Delete(); err != nil {
		return err
	}
	return nil
}
