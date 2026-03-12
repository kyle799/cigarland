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

func HandleUpdateCigar(ctx *gin.Context) {
	var cigar Cigar
	if err := ctx.ShouldBindJSON(&cigar); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if result := cigarDB.Save(&cigar); result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(200, cigar)
}

func HandleGetWrappers(ctx *gin.Context) {
	var wrappers []string
	cigarDB.Model(&Cigar{}).Distinct("\"Wrapper\"").Where("\"Wrapper\" != ''").Pluck("Wrapper", &wrappers)
	ctx.JSON(http.StatusOK, wrappers)
}
