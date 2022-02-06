package app

import (
	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

func StartApplication() {
	urls()
	r.Run(":8080")
}
