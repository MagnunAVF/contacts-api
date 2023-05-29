package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetContactsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"route": "get contacts",
	})
}

func GetContactHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"route": "get contact",
		"id":    id,
	})
}

func CreateContactHandler(c *gin.Context) {
	var contact Contact
	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"route":   "create contact",
		"contact": contact,
	})
}

func UpdateContactHandler(c *gin.Context) {
	id := c.Param("id")

	var contact Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"route":   "create contact",
		"contact": contact,
		"id":      id,
	})
}

func DeleteContactHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"route": "delete contact",
		"id":    id,
	})
}
