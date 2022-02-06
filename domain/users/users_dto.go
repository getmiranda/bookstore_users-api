package users

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID          uint64    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"date_created"`
}

func (u *User) Validate() error {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.New("invalid email address")
	}
	return nil
}
