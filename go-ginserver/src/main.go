package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	app.GET("health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
	app.Run("0.0.0.0:8080")
}
