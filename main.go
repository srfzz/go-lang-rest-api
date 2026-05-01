package main

import (
	"log/slog"
	"net/http"

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
	server.Run(":8080")
}

func getAllEvents(context *gin.Context) {
	events := models.GetAllEvents()
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
	event.UserId = 1
	event.ID = 1
	event.Save()
	contxet.JSON(http.StatusCreated, gin.H{"message": "Event created !", "event": event})
}
