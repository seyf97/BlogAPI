package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seyf97/BlogAPI/models"
)

func getAllArticles(context *gin.Context) {
	articles, err := models.GetAllArticlesDB()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch articles.", "err_msg": err.Error()})
		return
	}

	context.JSON(http.StatusOK, articles)
}

func getSingleArticle(context *gin.Context) {
	article_id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse article id."})
		return
	}

	article, err := models.GetArticleByIdDB(article_id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch article."})
		return
	}

	context.JSON(http.StatusOK, *article)
}

func createArticle(context *gin.Context) {
	var article models.Article
	err := context.ShouldBindJSON(&article)

	// userID saved at middleware level
	article.User_ID = context.GetInt64("userID")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Set the CreatedAt to the current time
	article.CreatedAt = time.Now()

	err = article.SaveDB()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not post article."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Posted article successfully."})
}

func deleteArticle(context *gin.Context) {
	article_id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse article id."})
		return
	}

	db_article, err := models.GetArticleByIdDB(article_id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find article."})
		return
	}

	incoming_user_ID := context.GetInt64("userID")
	if db_article.User_ID != incoming_user_ID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Only authors can delete articles."})
		return
	}

	err = db_article.DeleteDB()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete article."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully."})
}

func updateArticle(context *gin.Context) {
	article_id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse article id."})
		return
	}

	db_article, err := models.GetArticleByIdDB(article_id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find article."})
		return
	}

	incoming_user_ID := context.GetInt64("userID")
	if incoming_user_ID != db_article.User_ID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Only authors can update articles."})
		return
	}

	var updated_article models.Article

	err = context.ShouldBindJSON(&updated_article)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Set last updated time and ID
	now := time.Now()
	updated_article.LastEdited = &now
	updated_article.ID = article_id

	err = updated_article.UpdateDB()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update article."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Article updated successfully."})
}
