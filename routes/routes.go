package routes

import (
	"github.com/gin-gonic/gin"
	"blog/handlers"
)

// SetupRoutes configures all application routes
func SetupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	ginServer := gin.Default()

	ginServer.SetTrustedProxies([]string{"100.20.92.101", "44.225.181.72", "44.227.217.144"})

	ginServer.LoadHTMLGlob("templates/*")
	ginServer.Static("/static", "./static")

	// Home page
	ginServer.GET("/", handlers.HomeHandler)

	// Health check
	ginServer.GET("/ping", handlers.PingResponse)

	// API routes
	apiGroup := ginServer.Group("/api/v1")
	{
		apiGroup.GET("/blogs", handlers.GetBlogs)
		apiGroup.GET("/blog/:id", handlers.GetBlogByID)
		apiGroup.GET("/movies", handlers.GetMovies)
	}

	// 404 handler
	ginServer.NoRoute(handlers.NotFoundHandler)

	return ginServer
}