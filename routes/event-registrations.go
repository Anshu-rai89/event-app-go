package routes

import (
	"net/http"
	"strconv"

	"github.com/Anshu-rai89/event-app-go/models"
	"github.com/gin-gonic/gin"
)

func registerEventForUser(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid eventId"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid eventId"})
		return
	}

	err = event.RegisterEvent(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Event registration success"})

}

func cancelEventRegistrationForUser(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid eventId"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid eventId"})
		return
	}

	err = event.CancelEventRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Event cancellation success"})
}
