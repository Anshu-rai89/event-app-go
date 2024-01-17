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

	userId := c.GetInt64("userId")
	event.UserId = userId
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

func updateEvent(c *gin.Context) {
	var event models.Event
	var eventId int64
	var err error
	eventId, err = strconv.ParseInt(c.Param("id"), 10, 64)
	userId := c.GetInt64("userId")

	storedEvent, err := models.GetEventByID(eventId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid eventId"})
		return
	}

	if storedEvent.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "UnAuthorized"})
	}

	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid data"})
		return
	}
	event.Id = eventId
	err = event.UpdateEvent()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Event updated success"})
}

func deleteEvent(c *gin.Context) {
	var eventId int64
	var err error
	eventId, err = strconv.ParseInt(c.Param("id"), 10, 64)
	userId := c.GetInt64("userId")

	storedEvent, err := models.GetEventByID(eventId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid eventId"})
		return
	}

	if storedEvent.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "UnAuthorized"})
	}

	err = models.DeleteEvent(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Event deleted success"})
}
