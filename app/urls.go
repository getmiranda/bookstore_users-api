package app

import (
	"github.com/getmiranda/bookstore_users-api/controllers/ping"
	"github.com/getmiranda/bookstore_users-api/controllers/users"
)

func urls() {
	r.GET("/ping", ping.Ping)

	r.GET("/users/:user_id", users.Get)
	r.PUT("/users/:user_id", users.Update)
	r.PATCH("/users/:user_id", users.Update)
	r.DELETE("/users/:user_id", users.Delete)
	r.GET("/internal/users/search", users.Search)

	// Private endpoints
	r.POST("/users", users.Create)
	r.POST("/users/login", users.Login)
}
