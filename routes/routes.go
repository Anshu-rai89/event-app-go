package routes

import (
	"net/http"

	"github.com/Anshu-rai89/event-app-go/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Hello world"})
	})

	authenticated := app.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/event", createEvent)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.POST("/event/:id/register", registerEventForUser)
	authenticated.DELETE("/event/:id/register", cancelEventRegistrationForUser)

	app.POST("/user/signup", createUser)
	app.POST("/user/login", loginUser)
	app.GET("/event", getEvents)
	app.GET("/event/:id", getEvent)
}
