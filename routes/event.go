package routes

import (
	"net/http"
	"strconv"

	"github.com/Anshu-rai89/event-app-go/models"
	"github.com/gin-gonic/gin"
)

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

func getEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid id"})
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Something went wrong."})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Event fetch success", "data": event})
}
