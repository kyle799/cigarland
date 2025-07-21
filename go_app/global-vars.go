package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

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

func CreateAndSendCigarCreationPayload(url *url.URL) {
	client := http.Client{}
	requestBody := &bytes.Buffer{}
	cigarPayload := CreateRandomCigarPayload()
	jsonContent, err := json.MarshalIndent(cigarPayload, " ", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling payload to json: %s", err)
		return
	}
	requestBody.Write(jsonContent)
	log.Printf("Content length of request: %d", requestBody.Len())
	fmt.Fprintf(os.Stderr, "%s", requestBody.Bytes())
	resp, err := client.Post(url.String(), "application/json", requestBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received from the server: %s", err)
	}
	fmt.Fprintf(os.Stdout, "Resopnse: %+v", resp)
}

func CreateRandomCigarPayload() *CigarCreatePayload {
	cigarPayload := &CigarCreatePayload{}
	cigarPayload.Cigars = CreateRandomCigars(10)
	return cigarPayload
}

func CreateRandomCigars(count int) []*Cigar {
	cigars := make([]*Cigar, count)
	for i := 0; i < count; i++ {
		cigars[i] = CreateRandomCigarPtr()
	}
	return cigars
}

func CreateRandomCigarPtr() *Cigar {
	c := &Cigar{}
	c.Brand = GenerateRandomAlnumString()
	c.Name = GenerateRandomAlnumString()
	c.Wrapper = GenerateRandomAlnumString()
	c.Profile = GenerateRandomAlnumString()
	c.TastyTip = GenerateRandomBool()
	c.Pressed = GenerateRandomBool()
	c.Spicy, _ = GenerateRandomIntInRange(0, 10)
	c.Rating, _ = GenerateRandomIntInRange(0, 10)
	c.Length, _ = GenerateRandomIntInRange(30, 90)
	c.Ring, _ = GenerateRandomIntInRange(40, 70)
	c.Review = GenerateRandomAlnumStringN(256)
	c.JohnRating, _ = GenerateRandomIntInRange(0, 10)
	c.KyleRating, _ = GenerateRandomIntInRange(0, 10)
	c.JohnReview = GenerateRandomAlnumStringN(256)
	c.KyleReview = GenerateRandomAlnumStringN(256)
	c.ImageRef = GenerateRandomAlnumStringN(128)
	c.AuthenticHumanReview = GenerateRandomAlnumStringN(256)
	return c
}
