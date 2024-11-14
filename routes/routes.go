package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/seyf97/BlogAPI/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	// Articles
	server.GET("/articles/:id", getSingleArticle)
	server.GET("/articles", getAllArticles)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/articles", createArticle)
	authenticated.DELETE("/articles/:id", deleteArticle)
	authenticated.PUT("/articles/:id", updateArticle)

	// Users
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
