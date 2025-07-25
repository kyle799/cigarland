package main

import (
	"gorm.io/gorm"
)

var (
	createDB          bool
	startServer       bool
	server            string
	testCigarCreation bool
	port              int
	dbPath            string
	dbName            string
	cigarDB           *gorm.DB
	apiPrefix         string
)

func CreateTableSchemas() []any {
	tableSchemas := make([]any, 0, 1)
	// tableSchemas = append(tableSchemas, CreateNewCigarTable())
	tableSchemas = append(tableSchemas, &Cigar{})
	return tableSchemas
}

func CreateNewCigarTable() *Cigar {
	return &Cigar{
		Brand:                "Default",
		Name:                 "Default",
		Wrapper:              "Default",
		Profile:              "String",
		TastyTip:             false,
		Binder:               "Default",
		Spicy:                0,
		Rating:               0,
		Length:               54,
		Ring:                 60,
		Review:               "",
		JohnRating:           0,
		JohnReview:           "",
		KyleRating:           0,
		KyleReview:           "",
		AuthenticHumanReview: "",
	}
}
