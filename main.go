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

  

  ginServer.GET("/ping", func(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  ginServer.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}