package app

import (
	"github.com/getmiranda/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

func StartApplication() {
	urls()
	logger.Info("starting application...")
	r.Run(":8080")
}
