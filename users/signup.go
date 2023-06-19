package users

import (
	"backend/helpers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = nil
	originalPassword := user.Password
	user.Password = helpers.Sha256(user.Password)

	if len(user.Username) < 4 || len(user.Password) < 4 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password must be longer than 4 charachters"})
		return
	}
	if user.Pfp == "" {
		user.Pfp = `https://i.pinimg.com/550x/57/70/f0/5770f01a32c3c53e90ecda61483ccb08.jpg`
	}

	db, dbErr := helpers.OpenDb()
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}
	result := db.Table("users").Create(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Unable to create user. Perhaps the username is taken?"})
		return
	}
	user.Password = originalPassword

	token, jwtErr := createJWT(user)
	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a JWT"})
		return
	}

	c.Header("Authorization", token)
	//fmt.Println(validateJWT(token))

	c.JSON(http.StatusOK, user)
}
