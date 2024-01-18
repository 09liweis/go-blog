package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Todo: add favicon

func requestHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("session", "user1")
		context.Next()

		//context.Abort()
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URL")

	if uri == "" {
		log.Fatal("You must set your 'MONGO_URL' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		// panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
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
			"msg": msg,
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

	apiGroup := ginServer.Group("/api")
	{
		apiGroup.GET("/blogs", func(context *gin.Context) {
			page := context.Query("page")
			limit := context.Query("limit")
			context.JSON(http.StatusOK, gin.H{
				"page":  page,
				"limit": limit,
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
