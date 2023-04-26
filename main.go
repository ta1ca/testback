package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	doMain()
}

func doMain() {
	r := gin.Default()

	// CORS設定
	r.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	})

	r.GET("/api/go", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"mmm": "hello world",
		})
	})
	r.Run(":8000")
}
