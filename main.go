package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	dotenv"github.com/joho/godotenv"

	"backend/helpers"
	"backend/middleware"
	"backend/posts"
	"backend/users"
	"backend/likes"
)

func main() {
	dotenv.Load(".env")
	port := os.Getenv("PORT")
	helpers.MakeMigrations()

	r := gin.Default()
	r.POST("/signup", users.Signup)
	r.POST("/signin", users.Signin)
	r.POST("/post", middleware.AuthMiddleware, posts.CreatePost)
	r.POST("like", middleware.AuthMiddleware, likes.CreateLike)
	r.DELETE("/post", middleware.AuthMiddleware, posts.DeletePost)
	r.DELETE("/like", middleware.AuthMiddleware, likes.DeleteLike)
	r.GET("/post", posts.GetPosts)
	r.GET("/comments", posts.GetPostComments)
	r.GET("/users", users.GetUser)

	fmt.Printf("serving on http://localhost:%s\n", port)
	r.Run()
}
