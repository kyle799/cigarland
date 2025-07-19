package main

import (
	"encoding/json"
	"fmt"
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
	bodyContent := make([]byte, ctx.Request.ContentLength)
	ctx.Request.Body.Read(bodyContent)
	cigarCreatePayload := CigarCreatePayload{}
	_ = json.Unmarshal(bodyContent, &cigarCreatePayload)
	cigarInfo := cigarCreatePayload.CigarInfo
	db.Create(cigarInfo)
}
