package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// "gorm.io/gorm/schema"

type NameKeeper struct{}

func (nk NameKeeper) Replace(field string) string {
	return field
}

type CustomNamingStrategy struct{}

func OpenDB(dbName string) *gorm.DB {
	fmt.Printf("opening db: %s\n", dbName)
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})
	if err != nil {
		log.Fatalf("Error opening db: %s", err)
	}
	return db
}

func InitializeDBTables(db *gorm.DB, tableSchemas []any) {
	log.Printf("Starting table migration\n")
	for _, schema := range tableSchemas {
		db.AutoMigrate(schema)
	}
}

func HandleCreateCigar(ctx *gin.Context, db *gorm.DB) {
	log.Printf("receiving request to insert cigar")
	log.Printf("making byte slice from requst content length: %d", ctx.Request.ContentLength)
	// bodyContent := make([]byte, ctx.Request.ContentLength)
	log.Printf("reading request body into byte slice")
	bodyContent, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Printf("Error reading from request body: %s", err)
	}
	// log.Printf("Read bytes from requeset body: %d", readBytes)
	log.Printf("Request Body: %s", bodyContent)
	log.Printf("creating cigar payload struct")
	cigarCreatePayload := CigarCreatePayload{}
	log.Printf("unmarshaling payload into cigarpayload struct")
	err = json.Unmarshal(bodyContent, &cigarCreatePayload)
	if err != nil {
		log.Printf("error unmarshaling json into cigarCreatePayload struct: %s", err)
	} else {
		log.Printf("CigarCreatePayload: %+v", cigarCreatePayload)
	}
	cigarInfo := cigarCreatePayload.Cigars
	for _, cigar := range cigarInfo {
		log.Printf("Creating cigar: %+v", *cigar)
		db.Create(cigar)
	}
}
