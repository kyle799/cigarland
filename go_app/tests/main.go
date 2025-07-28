package main

import (
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
	output             string
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
	if dumpJson {
		outputWriter := ParseOutputValue()
		DumpToJson(&Cigar{}, outputWriter)
	}
}

func RegisterAndParseCLIArgs() {
	flag.BoolVar(&testCigarInsertion, "test-cigar-insertion", false, "Boolean toggle for testing inserting cigars into the DB via the api")
	flag.BoolVar(&testCigarQuery, "query", false, "Boolean toggle for testing a SELECT * from the cigar table")
	flag.BoolVar(&dumpJson, "dump", false, "Boolean toggle to dump json schema for objects")
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
