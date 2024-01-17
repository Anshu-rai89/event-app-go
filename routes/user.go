package routes

import (
	"fmt"
	"net/http"

	"github.com/Anshu-rai89/event-app-go/models"
	"github.com/Anshu-rai89/event-app-go/utils"
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
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "User created success"})
}

func loginUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid data"})
		return
	}

	err = user.ValidatePassword()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	fmt.Println("token", token, err)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Login Success", "data": token})
}
