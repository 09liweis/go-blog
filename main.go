package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

//Todo: add favicon

func main() {
  ginServer := gin.Default()


  ginServer.LoadHTMLGlob("templates/*")
  ginServer.Static("/static", "./static")

  ginServer.GET("/", func(context *gin.Context) {
    context.HTML(http.StatusOK, "index.html",gin.H{
      "msg":"first msg",
    })
  })

  ginServer.NoRoute(func(context *gin.Context) {
    context.HTML(http.StatusNotFound, "404.html",nil)
  })

  ginServer.GET("/ping", func(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  apiGroup := ginServer.Group("/api")
  {
    apiGroup.GET("/blogs", func(context *gin.Context)  {
      page := context.Query("page")
      limit := context.Query("limit")
      context.JSON(http.StatusOK, gin.H{
        "page": page,
        "limit": limit,
      })
    })

    apiGroup.GET("/blog/:id", func(context *gin.Context)  {
      blogId := context.Param("id")
      context.JSON(http.StatusOK, gin.H{
        "blogId": blogId,
      })
    })
  }



  ginServer.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}