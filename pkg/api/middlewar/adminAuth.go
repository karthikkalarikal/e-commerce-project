package middlewar

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AdminMiddleware(c *gin.Context) {
	fmt.Println("*****Middle ware Auth******8")
	tokenString := c.GetHeader("Authorization")
	// fmt.Println("here")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("super-secret-key"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid autorization token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid autorization token"})
		c.Abort()
		return
	}

	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		c.Abort()
		return
	}

	c.Set("role", role)

	c.Next()
}
