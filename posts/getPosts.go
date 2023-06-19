package posts

import (
	"backend/helpers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	db, dbError := helpers.OpenDb()
	if dbError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbError.Error()})
		return
	}

	var posts []models.Post
	result := db.Table("posts").Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func GetPostComments(c *gin.Context) {
	id := c.Query("id")
	if id == "" {c.JSON(http.StatusBadRequest, gin.H{"error":"An id field is needed (url/comments?id=10) for example"}); return}
	db, dbError := helpers.OpenDb()
	if dbError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbError.Error()})
		return
	}
	var comments []models.Post
	result := db.Table("posts").Where("under = ?", id).Find(&comments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)

}