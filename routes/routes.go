package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getAllEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:eventId", getEvent)
}
