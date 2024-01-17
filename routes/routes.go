package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Hello world"})
	})
	app.POST("/event", createEvent)
	app.GET("/event", getEvents)
	app.GET("/event/:id", getEvent)
	app.PUT("/event/:id", updateEvent)
	app.POST("/user/signup", createUser)
}
