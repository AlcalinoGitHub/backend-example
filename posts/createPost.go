package posts

import (
	"backend/helpers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	post := models.Post{}
	err := c.Bind(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if post.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post content empty or mussing"})
		return
	}
	post.ID = nil
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
	post.Author = *userObj.ID

	db, dbErr := helpers.OpenDb()
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}
	result := db.Table("posts").Create(&post)
	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}


	c.JSON(http.StatusOK, post)
}
