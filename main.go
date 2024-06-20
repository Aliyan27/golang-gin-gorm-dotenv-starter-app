package main

import (
	"ProjectStarter/config"
	"ProjectStarter/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadDotEnv()
	router := gin.Default()
	routes.Routes(router)
	router.Run(config.AppConfig.Port)
}