package users

import (
	"backend/helpers"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		id = 0
	}
	db, dbErr := helpers.OpenDb()
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}
	var users []models.User
	result := db.Table("users").Where("id = ? OR ? = 0", id, id).Find(&users)
	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldnt find any users"})
		return
	}
	for i := 0; i < len(users); i ++ {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, users)
}