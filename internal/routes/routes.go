package routes

import (
	"github.com/gin-gonic/gin"
)

func RoutesHandler() {
	r := gin.Default()
	AuthRoutes(r)
	r.Run(":8080")
}
