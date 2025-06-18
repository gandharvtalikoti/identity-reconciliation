package routes

import (
	"bitespeed-identity/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/identify", controllers.IdentifyContact)
	return r
}
