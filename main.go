package main

import (
	"net/http"

	"github.com/Anshu-rai89/event-app-go/db"
	"github.com/Anshu-rai89/event-app-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Hello world"})
	})

	app.POST("/event", createEvent)
	app.GET("/event", getEvents)
	app.Run()
}

func createEvent(c *gin.Context) {
	db.InitDB()
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid data"})
		return
	}
	event.Id = 1
	event.UserId = 1
	event.Save()
	c.JSON(http.StatusOK, gin.H{"msg": "Event created success"})
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, gin.H{"msg": "Events fetch success", "data": events})
}
