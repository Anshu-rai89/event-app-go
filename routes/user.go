package routes

import (
	"net/http"

	"github.com/Anshu-rai89/event-app-go/models"
	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid data"})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "User created success"})
}
