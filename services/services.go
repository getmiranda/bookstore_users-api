package services

import (
	"github.com/getmiranda/bookstore_users-api/domain/users"
	"github.com/getmiranda/bookstore_users-api/services/users_service"
)

var UsersService usersService

type usersService interface {
	GetUser(uint64) (*users.User, error)
	CreateUser(*users.User) (*users.User, error)
	UpdateUser(bool, *users.User) (*users.User, error)
	DeleteUser(uint64) error
	Search(string) (users.Users, error)
	LoginUser(*users.LoginRequest) (*users.User, error)
}

func init() {
	UsersService = &users_service.UsersService{}
}
