package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"blog/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Blog represents a blog post structure
type Blog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	Author    string             `json:"author" bson:"author"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// PingResponse handles ping requests
func PingResponse(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// GetBlogs handles getting blogs list
func GetBlogs(context *gin.Context) {
	client := database.GetMongoClient()
	if client == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	// Parse pagination parameters
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))
	
	// Calculate skip value for pagination
	skip := (page - 1) * limit

	// Set up options for pagination and sorting
	opts := options.Find()
	opts.SetSkip(int64(skip))
	opts.SetLimit(int64(limit))
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}}) // Sort by creation date, newest first

	// Find blogs
	cursor, err := client.Database("heroku_6njptcbp").Collection("blogs").Find(context.Request.Context(), bson.D{}, opts)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Request.Context())

	// Map results
	var blogs []Blog
	if err = cursor.All(context.Request.Context(), &blogs); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get total count for pagination info
	total, err := client.Database("heroku_6njptcbp").Collection("blogs").CountDocuments(context.Request.Context(), bson.D{})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return response with pagination info
	context.JSON(http.StatusOK, gin.H{
		"blogs": blogs,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": (total + int64(limit) - 1) / int64(limit), // Calculate total pages
		},
	})
}

// GetBlogByID handles getting a single blog by ID
func GetBlogByID(context *gin.Context) {
	client := database.GetMongoClient()
	if client == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	blogId := context.Param("id")
	if blogId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Blog ID is required"})
		return
	}

	var blog Blog
	var err error

	// First try to convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(blogId)
	if err == nil {
		// Try to find by ObjectID
		err = client.Database("heroku_6njptcbp").Collection("blogs").FindOne(context.Request.Context(), bson.M{"_id": objectID}).Decode(&blog)
	} else {
		// If not a valid ObjectID, try to find by string ID
		err = client.Database("heroku_6njptcbp").Collection("blogs").FindOne(context.Request.Context(), bson.M{"_id": blogId}).Decode(&blog)
	}

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			context.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	context.JSON(http.StatusOK, blog)
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