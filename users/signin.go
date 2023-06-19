package users

import (
	"backend/helpers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signin(c *gin.Context) {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, dbErr := helpers.OpenDb()
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}

	user.Password = helpers.Sha256(user.Password)
	user.ID = nil
	result := db.Table("users").Where("username = ? AND password = ?", user.Username, user.Password).First(&user)
	if result.RowsAffected == 0 || result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not find a user"})
		return
	}

	token, jwtErr := createJWT(user)
	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a JWT"})
		return
	}
	c.Header("Authorization", token)

	c.JSON(http.StatusOK, user)

}
