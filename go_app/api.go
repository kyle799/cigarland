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

func HandleGetWrappers(ctx *gin.Context) {
	var wrappers []string
	cigarDB.Model(&Cigar{}).Distinct("wrapper").Where("wrapper != ''").Pluck("wrapper", &wrappers)
	ctx.JSON(http.StatusOK, wrappers)
}
