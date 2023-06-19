package likes

import (
	"backend/helpers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteLike(c *gin.Context) {
	like := models.Like{}
	err := c.Bind(&like)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get user data"})
		return
	}
	userObj, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not parse user data"})
		return
	}
	db, dbErr := helpers.OpenDb()
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}

	result := db.Table("likes").Where("under = ? AND author = ?", like.Under, userObj.ID).Delete(nil)
	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete like"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Like deleted succesfully"})
}
