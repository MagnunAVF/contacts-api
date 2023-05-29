package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func init() {
	log.Printf("Starting Gin ...")

	R = gin.Default()

	R.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
