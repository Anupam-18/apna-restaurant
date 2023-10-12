package middleware

import (
	"apna-restaurant/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.GetHeader("Authorization")
		headerToken = strings.TrimSpace(headerToken)
		if len(headerToken) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing Authorization"})
			c.Abort()
			return
		}
		splittedToken := strings.Split(headerToken, " ")
		if len(splittedToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Malformed authorization token"})
			c.Abort()
			return
		}
		tokenObj := &models.Token{}
		token, err := jwt.ParseWithClaims(splittedToken[1], tokenObj, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Malformed authorization token"})
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden, invalid token"})
			c.Abort()
			return
		}
		if tokenObj.ExpiresAt < time.Now().Local().Unix() {
			c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden, expired token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
