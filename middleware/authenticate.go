package middleware

import (
	"backend/models"
	"backend/users"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	claims, err := users.ValidateJWT(authHeader)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	var user models.User
	decodeErr := mapstructure.Decode(claims, &user)
	if decodeErr != nil {
		c.AbortWithError(http.StatusUnauthorized, decodeErr)
	}
	c.Set("user", user)
	c.Next()
}
