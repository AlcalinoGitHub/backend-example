package likes

import (
	"backend/helpers"
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLike(c *gin.Context) {
	like := models.Like{}
	err := c.Bind(&like)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if like.Under == 0 {c.JSON(http.StatusBadRequest, gin.H{"error": "A post id (under) is required"});return}
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
	like.ID = nil
	like.Author = *userObj.ID
	like.Matcher = fmt.Sprintf("%d => %d", like.Author, like.Under)
	db, dbErr := helpers.OpenDb()
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}
	result := db.Table("likes").Create(&like)
	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "A user cannot like the same post twice"})
		return
	}
	c.JSON(http.StatusOK, like)
}