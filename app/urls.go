package app

import (
	"github.com/getmiranda/bookstore_users-api/controllers/ping"
	"github.com/getmiranda/bookstore_users-api/controllers/users"
)

func urls() {
	r.GET("/ping", ping.Ping)

	r.GET("/users/:user_id", users.GetUser)
	r.POST("/users", users.CreateUser)
}
