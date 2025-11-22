package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"blog/database"
)

// PingResponse handles ping requests
func PingResponse(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// GetBlogs handles getting blogs list
func GetBlogs(context *gin.Context) {
	page := context.Query("page")
	limit := context.Query("limit")
	context.JSON(http.StatusOK, gin.H{
		"page":  page,
		"limit": limit,
		"array": []string{"a", "b", "c"},
	})
}

// GetBlogByID handles getting a single blog by ID
func GetBlogByID(context *gin.Context) {
	blogId := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"id": blogId,
	})
}

// GetMovies handles getting movies from MongoDB
func GetMovies(c *gin.Context) {
	client := database.GetMongoClient()
	if client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	// Find movies
	cursor, err := client.Database("heroku_6njptcbp").Collection("visuals").Find(c.Request.Context(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map results
	var movies []gin.H
	if err = cursor.All(c.Request.Context(), &movies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return movies
	c.JSON(http.StatusOK, movies)
}

// HomeHandler handles the home page
func HomeHandler(context *gin.Context) {
	const msg = "Hello World"
	context.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Go Blogs",
		"msg":   msg,
	})
}

// NotFoundHandler handles 404 errors
func NotFoundHandler(context *gin.Context) {
	context.HTML(http.StatusNotFound, "404.html", gin.H{
		"title": "Go Blogs 404",
	})
}