package routes

import (
	"github.com/gin-gonic/gin"
	"sep.com/eventapi/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)   //set handler for get request
	server.GET("events/:id", getEvent) //gin enable  developer to make dynamic endpoint with : + random name

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
