package routes

import (
	"github.com/gin-gonic/gin"
	"gitub.com/Jidetireni/events-restapi/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	auth := server.Group("/")
	auth.Use(middleware.Auth)
	auth.POST("/events", createEvent)
	auth.PUT("/events/:id", updateEvent)
	auth.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
