package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var APP *gin.Engine

func getHomePageHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "Welcome!"})
}

func main() {
	APP = gin.Default()
	APP.GET("/", getHomePageHandler)
	if os.Getenv("LCP") == "LOCAL" {
		APP.Run("localhost:8081")
	} else {
		port := os.Getenv("PORT")
		APP.Run(":" + port)
	}
}
