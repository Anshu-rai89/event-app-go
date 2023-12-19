package main

import (
	"net/http"

	"github.com/Anshu-rai89/event-app-go/db"
	"github.com/Anshu-rai89/event-app-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	app := gin.Default()
	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Hello world"})
	})

	app.POST("/event", createEvent)
	app.GET("/event", getEvents)
	app.Run()
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid data"})
		return
	}
	event.Id = 1
	event.UserId = 1
	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Event created success"})
}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Events fetch success", "data": events})
}
