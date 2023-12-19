package main

import (
	"github.com/Anshu-rai89/event-app-go/db"
	"github.com/Anshu-rai89/event-app-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	app := gin.Default()
	routes.RegisterRoutes(app)
	app.Run()
}
