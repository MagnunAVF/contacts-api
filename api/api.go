package api

import (
	"log"

	"github.com/MagnunAVF/contacts-api/api/handlers"
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func init() {
	log.Printf("Starting Gin ...")

	R = gin.Default()

	R.GET("/health", handlers.HealthHandler)

	R.GET("/contacts", handlers.GetContactsHandler)
	R.GET("/contacts/:id", handlers.GetContactHandler)
	R.POST("/contacts", handlers.CreateContactHandler)
	R.PUT("/contacts/:id", handlers.UpdateContactHandler)
	R.DELETE("/contacts/:id", handlers.DeleteContactHandler)
}
