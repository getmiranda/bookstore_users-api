package app

import (
	"github.com/getmiranda/bookstore_users-api/controllers/ping"
	"github.com/getmiranda/bookstore_users-api/controllers/users"
)

func urls() {
	r.GET("/ping", ping.Ping)

	r.GET("/users/:user_id", users.Get)
	r.POST("/users", users.Create)
	r.PUT("/users/:user_id", users.Update)
	r.PATCH("/users/:user_id", users.Update)
	r.DELETE("/users/:user_id", users.Delete)
}
