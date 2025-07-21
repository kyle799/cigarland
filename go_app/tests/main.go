package main

import (
	"flag"
	"log"
	"net/url"
	"os"
)

var (
	testCigarInsertion bool
	apiURL             string
	lInfo              = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	lErr               = log.New(os.Stderr, "ERROR: ", log.LstdFlags)
	testCigarQuery     bool

	// port               int
)

func main() {
	RegisterAndParseCLIArgs()
	if testCigarInsertion {
		testURL, err := url.Parse(apiURL)
		if err != nil {
			lErr.Printf("Error parsing URL: %s\n", err)
		}
		CreateAndSendCigarCreationPayload(testURL)
	}
}

func RegisterAndParseCLIArgs() {
	flag.BoolVar(&testCigarInsertion, "test-cigar-insertion", true, "Boolean toggle for testing inserting cigars into the DB via the api")
	flag.BoolVar(&testCigarQuery, "query", false, "Boolean toggle for testing a SELECT * from the cigar table")
	flag.StringVar(&apiURL, "api-url", "http://localhost:8080/api/test", "Url to use for testing the api")
	flag.StringVar(&apiURL, "host", "http://localhost:8080/api/test", "Url to use for testing the api")
	// flag.IntVar(&port, "port", 8080, "Port to use for accessing the api")
	flag.Parse()
}
