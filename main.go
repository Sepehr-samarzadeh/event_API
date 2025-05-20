package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sep.com/eventapi/db"
	"sep.com/eventapi/models"
)

func main() {
	db.InitDB()

	server := gin.Default()
	server.GET("/events", getEvents)   //set handler for get request
	server.GET("events/:id", getEvent) //gin enable  developer to make dynamic endpoint with : + random name
	server.POST("/events", createEvent)
	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, event)
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
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error saving events"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created"})
}
