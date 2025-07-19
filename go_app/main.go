package main

import (
	// "os"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createdbfile() {
	db, err := sql.Open("sqlite3", "/cigarland/db/cigar.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test successful",
	})
}

func main() {
	var server string = "localhost:8080"
	fmt.Printf("starting webserver at %s", server)
	router := gin.Default()
	router.GET("/test", test)
	router.Run(server)
}
