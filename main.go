package main

import (
	handler "sizzlie/handler/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/signup", func(c *gin.Context) {
		handler.Create(c)
	})

	router.Run(":8000")
}
