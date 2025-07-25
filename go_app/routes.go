package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func SetRoutesAndRun(router *gin.Engine) {
	router.GET(fmt.Sprintf("%s/test", apiPrefix), test)
	router.POST(fmt.Sprintf("%s/test", apiPrefix), HandleCreateCigarRouter)
	router.GET(fmt.Sprintf("%s/cigars", apiPrefix), HandleQueryCigarTable)
	log.Printf("starting webserver at %s:%d", server, port)
	router.Run(fmt.Sprintf("%s:%d", server, port))
}
