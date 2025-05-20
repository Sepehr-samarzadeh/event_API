package main

import (
	"github.com/gin-gonic/gin"
	"sep.com/eventapi/db"
	"sep.com/eventapi/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")

}
