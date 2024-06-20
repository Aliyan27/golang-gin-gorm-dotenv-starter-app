package routes

import (
	"ProjectStarter/controller"
	"ProjectStarter/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	api := router.Group("/api").Use(middleware.AuthMiddleware())
	api.GET("/",controller.HelloWorld)
}