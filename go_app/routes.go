package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func SetRoutesAndRun(router *gin.Engine) {
	// Auth routes (public)
	router.GET("/login", HandleLogin)
	router.GET("/logout", HandleLogout)
	router.GET("/auth/google/callback", HandleOAuthCallback)

	// API routes (protected)
	api := router.Group(apiPrefix, WithAuth())
	api.GET("/test", test)
	api.POST("/test", HandleCreateCigarRouter)
	api.GET("/cigars", MakeDBHandler(cigarDB, HandleQueryCigarTable))
	api.POST("/where", MakeDBHandler(cigarDB, SelectWhereHandler))

	log.Printf("starting webserver at %s:%d", server, port)
	router.Run(fmt.Sprintf("%s:%d", server, port))
}
