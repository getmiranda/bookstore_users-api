package users

import (
	"encoding/json"
	"time"
)

type PublicUser struct {
	ID          uint64    `json:"id"`
	DateCreated time.Time `json:"date_created"`
	Status      string    `json:"status"`
}

type PrivateUser struct {
	ID          uint64    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"date_created"`
	Status      string    `json:"status"`
}

func (u Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(u))
	for index, user := range u {
		result[index] = user.Marshal(isPublic)
	}
	return result
}

func (u *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          u.ID,
			DateCreated: u.DateCreated,
			Status:      u.Status,
		}
	}
	userJson, _ := json.Marshal(u)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
	// return PrivateUser{
	// 	ID:          u.ID,
	// 	FirstName:   u.FirstName,
	// 	LastName:    u.LastName,
	// 	Email:       u.Email,
	// 	DateCreated: u.DateCreated,
	// 	Status:      u.Status,
	// }
}
