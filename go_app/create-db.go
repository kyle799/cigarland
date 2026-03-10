package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type NameKeeper struct{}

func (nk NameKeeper) Replace(field string) string {
	return field
}

func OpenDB(dsn string) *gorm.DB {
	fmt.Printf("opening db\n")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
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
