package main

import (
	"github.com/gin-gonic/gin"
	"gitub.com/Jidetireni/events-restapi/db"
	"gitub.com/Jidetireni/events-restapi/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
