package main

import (
	"appointment-tracking/db"
	"appointment-tracking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	if err := server.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
