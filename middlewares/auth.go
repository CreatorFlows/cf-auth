package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/cf-auth/utils"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("tokenCF")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(cookie)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
			c.Abort()
			return
		}

		c.Set("role", claims.Role)
		c.Next()
	}
}
