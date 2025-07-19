package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "test successful",
    })
}

func main() {
    router := gin.Default()
    router.GET("/test", test)
    router.Run("localhost:8080")
}
