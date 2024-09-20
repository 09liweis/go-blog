package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

//Todo: add favicon

func requestHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("session", "user1")
		context.Next()

		//context.Abort()
	}
}

var uri = os.Getenv("MONGODB_URL")
var mongoClient *mongo.Client

// This function runs before we call our main function and connects to our MongoDB database. If it cannot connect, the application stops.
func init() {
	if err := connect_to_mongodb(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
}

func getMovies(c *gin.Context) {
	// Find movies
	cursor, err := mongoClient.Database("heroku_6njptcbp").Collection("visuals").Find(context.TODO(), bson.D{{}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map results
	var movies []bson.M
	if err = cursor.All(context.TODO(), &movies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return movies
	c.JSON(http.StatusOK, movies)
}

// Our implementation code to connect to MongoDB at startup
func connect_to_mongodb() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	mongoClient = client
	return err
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if uri == "" {
		log.Fatal("You must set your 'MONGO_URL' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	// mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	// if err != nil {
	// 	panic(err)
	// }
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			// panic(err)
		}
	}()

	gin.SetMode(gin.ReleaseMode)

	ginServer := gin.Default()

	ginServer.SetTrustedProxies([]string{"100.20.92.101", "44.225.181.72", "44.227.217.144"})

	ginServer.LoadHTMLGlob("templates/*")
	ginServer.Static("/static", "./static")

	ginServer.GET("/", func(context *gin.Context) {
		const msg = "Hello World"
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Go Blogs",
			"msg":   msg,
		})
	})

	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	ginServer.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	apiGroup := ginServer.Group("/api/v1") {
		apiGroup.GET("/blogs", func(context *gin.Context) {
			page := context.Query("page")
			limit := context.Query("limit")
			context.JSON(http.StatusOK, gin.H{
				"page":  page,
				"limit": limit,
				"array": []string{"a", "b", "c"},
			})
		})

		apiGroup.GET("/blog/:id", func(context *gin.Context) {
			blogId := context.Param("id")
			context.JSON(http.StatusOK, gin.H{
				"id": blogId,
			})
		})
	}

	ginServer.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
