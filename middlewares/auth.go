package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Anshu-rai89/event-app-go/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	fmt.Println("TOken", token)
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
		return
	}

	c.Set("userId", userId)
	c.Next()
}
