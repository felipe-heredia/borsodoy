package middleware

import (
	"borsodoy/radovid/pkg/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		token, err := utility.ValidateToken(tokenString)

		if err != nil {
			if httpError, ok := err.(*utility.HttpError); ok {
				c.JSON(httpError.Status, gin.H{"message": httpError.Message})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error middleware"})
			return
		}

    c.Set("email", token.Cliams.Email)
		c.Next()
	}
}
