package users

import (
	"errors"
	"strings"
)

type User struct {
	ID          uint64 `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (u *User) Validate() error {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.New("invalid email address")
	}
	return nil
}
