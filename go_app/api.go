package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test successful",
	})
}

func HandleCreateCigarRouter(ctx *gin.Context) {
	log.Printf("Received request to create cigars")
	HandleCreateCigar(ctx, cigarDB)
}
