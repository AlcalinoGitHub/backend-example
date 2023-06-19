package posts

import (
	"backend/helpers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletePost(c *gin.Context) {
	post := models.Post{}
	err := c.Bind(&post)
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
	result := db.Table("posts").Where("id = ? AND author = ?", post.ID, userObj.ID).Delete(nil)
	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the post. Perhaps the ID is wrong?"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "post deleted succesfully"})

}
