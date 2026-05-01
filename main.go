package main

import (
	"log/slog"
	"net/http"
	"strconv"

	"go-lang-restapi/db"
	"go-lang-restapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	defer db.DB.Close()
	server := gin.Default()
	server.GET("/events", getAllEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:eventId", getEvent)
	server.Run(":8080")
}

func getAllEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		slog.Error(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong", "error": err.Error()})
	}
	slog.Info("/events Route Accessed ", " ip: ", context.ClientIP())
	context.JSON(http.StatusOK, events)
}

func createEvent(contxet *gin.Context) {
	var event models.Event
	err := contxet.ShouldBindJSON(&event)
	if err != nil {
		contxet.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "error": err.Error()})
		return
	}
	var userId uint64 = 1
	event.UserId = &userId
	err = event.Save()
	if err != nil {
		contxet.JSON(http.StatusInternalServerError, gin.H{"mesage ": "something went wrong", "errors": err.Error()})
		return
	}
	contxet.JSON(http.StatusCreated, gin.H{"message": "Event created !", "event": event})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Error fetching event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}
