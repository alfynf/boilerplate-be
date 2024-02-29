package routes

import (
	"boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", controllers.PingHandler)

	return r
}
