package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
)

var (
	testCigarInsertion bool
	apiURL             string
	dumpJson           bool
	debug              bool
	output             string
	lInfo              = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	lErr               = log.New(os.Stderr, "ERROR: ", log.LstdFlags)
	testCigarQuery     bool
	testCigarFilter    bool

	// port               int
)

func main() {
	RegisterAndParseCLIArgs()
	PopulateValueOperatorMap()
	encoder := json.NewEncoder(lInfo.Writer())
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	testURL, err := url.Parse(apiURL)
	if err != nil {
		log.Fatalf("Error parsing url: %s", err)
	}
	if testCigarInsertion {
		CreateAndSendCigarCreationPayload(testURL)
	}
	if testCigarQuery {
		// outputWriter := ParseOutputValue()
		queryURL, _ := url.Parse(apiURL)
		cigars := QueryCigarTable(queryURL)
		encoder.Encode(cigars)
		// DumpToJson(cigars, outputWriter)
	}
	if testCigarFilter {
		cigars, err := FilterCigarTableByRating(testURL)
		if err != nil {
			log.Fatalf("Error filtering cigars by rating: %s", err)
		}
		encoder.Encode(cigars)
	}
	if dumpJson {
		outputWriter := ParseOutputValue()
		DumpToJson(&Cigar{}, outputWriter)
	}
}

func RegisterAndParseCLIArgs() {
	flag.BoolVar(&testCigarInsertion, "test-cigar-insertion", false, "Boolean toggle for testing inserting cigars into the DB via the api")
	flag.BoolVar(&testCigarQuery, "query", false, "Boolean toggle for testing a SELECT * from the cigar table")
	flag.BoolVar(&dumpJson, "dump", false, "Boolean toggle to dump json schema for objects")
	flag.BoolVar(&testCigarFilter, "filter", false, "Boolean toggle for testing filtered queries")
	flag.BoolVar(&debug, "d", false, "Set debug mode")
	flag.BoolVar(&debug, "debug", false, "Set debug mode")
	flag.StringVar(&apiURL, "api-url", "http://localhost:8080/api/test", "Url to use for testing the api")
	flag.StringVar(&apiURL, "host", "http://localhost:8080/api/test", "Url to use for testing the api")
	flag.StringVar(&output, "output", "stdout", "Location for any output, defaults to stdout")
	// flag.IntVar(&port, "port", 8080, "Port to use for accessing the api")
	flag.Parse()
}

func ParseOutputValue() io.Writer {
	switch output {
	case "stdout":
		return os.Stdout
	case "stderr":
		return os.Stderr
	default:
		f, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening output file (%s): %s", output, err)
			return os.Stdout
		}
		return f
	}
}
