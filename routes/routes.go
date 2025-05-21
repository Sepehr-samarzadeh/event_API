package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)   //set handler for get request
	server.GET("events/:id", getEvent) //gin enable  developer to make dynamic endpoint with : + random name
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("events/:id", deleteEvent)
}
