package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sep.com/eventapi/db"
	"sep.com/eventapi/models"
)

func main() {
	db.InitDB()

	server := gin.Default()
	server.GET("/events", getEvents) //set handler for get request
	server.POST("/events", createEvent)
	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) { //when we use endpoint handler we are forced to use gin.context
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse data"})
		return
	}
	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "event created"})
}
