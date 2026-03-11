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
	router.GET("/me", HandleMe)

	// API routes
	api := router.Group(apiPrefix, WithAuth())
	api.GET("/test", test)
	api.GET("/cigars", MakeDBHandler(cigarDB, HandleQueryCigarTable))
	api.POST("/where", MakeDBHandler(cigarDB, SelectWhereHandler))
	api.POST("/test", WithPermission(func(p *UserPermission) bool { return p.CanAdd }), HandleCreateCigarRouter)
	api.DELETE("/cigars", WithPermission(func(p *UserPermission) bool { return p.CanDelete }), HandleDeleteCigar)
	api.GET("/admin/users", WithPermission(func(p *UserPermission) bool { return p.CanAdmin }), HandleAdminListUsers)
	api.POST("/admin/users", WithPermission(func(p *UserPermission) bool { return p.CanAdmin }), HandleAdminUpdateUser)

	log.Printf("starting webserver at %s:%d", server, port)
	router.Run(fmt.Sprintf("%s:%d", server, port))
}
