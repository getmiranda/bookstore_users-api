package users_service

import (
	"github.com/getmiranda/bookstore_users-api/domain/users"
	"github.com/getmiranda/bookstore_users-api/utils/crypto_utils"
	"github.com/getmiranda/bookstore_users-api/utils/date_utils"
	"github.com/getmiranda/bookstore_users-api/utils/errors"
)

type UsersService struct{}

func (u *UsersService) GetUser(userId uint64) (*users.User, error) {
	result := &users.User{ID: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UsersService) CreateUser(user *users.User) (*users.User, error) {
	if err := user.Validate(); err != nil {
		return nil, errors.NewBadRequestError("invalid parameters")
	}
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	user.DateCreated = date_utils.GetNow()

	if err := user.Save(); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UsersService) UpdateUser(isPartial bool, user *users.User) (*users.User, error) {
	current, err := u.GetUser(user.ID)
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

func (u *UsersService) DeleteUser(userId uint64) error {
	user := &users.User{ID: userId}
	return user.Delete()
}

func (u *UsersService) Search(status string) (users.Users, error) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (u *UsersService) LoginUser(request *users.LoginRequest) (*users.User, error) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
