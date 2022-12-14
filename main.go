package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lenichols/wallet-tag-api/config"
	"github.com/lenichols/wallet-tag-api/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}
