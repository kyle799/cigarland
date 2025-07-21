package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	//"sync"
	//"time"

	"github.com/gin-gonic/gin"
)

func createPath(path string) {
	err := os.MkdirAll(path, 0o750)
	if err != nil {
		log.Fatalf("failure to create %s", err)
	}
	log.Printf("Created path %s", path)
}

func main() {
	ParseArgs()
	if createDB {
		fmt.Printf("Starting creation of db\n")
		createPath(dbPath)
		tables := CreateTableSchemas()
		cigarDB = OpenDB(dbName)
		InitializeDBTables(cigarDB, tables)

	}
	if startServer {
		log.Printf("starting webserver at %s:%d", server, port)

		router := gin.Default()
		router.GET(fmt.Sprintf("%s/test", apiPrefix), test)
		router.POST(fmt.Sprintf("%s/test", apiPrefix), HandleCreateCigarRouter)
		router.Run(fmt.Sprintf("%s:%d", server, port))
	}
	if testCigarCreation && !startServer {
		url, err := url.Parse(fmt.Sprintf("http://localhost:%d/api/test", port))
		if err != nil {
			log.Printf("error creating url for localhost: %s\n", err)
			return
		} else {
			log.Printf("Hitting url with payload: %s\n", url.String())
			CreateAndSendCigarCreationPayload(url)
		}
	}
}

func ParseArgs() {
	flag.BoolVar(&createDB, "create-db", false, "Boolean toggle to create a new DB")
	flag.BoolVar(&startServer, "start-server", false, "Boolean toggle to start the server")
	flag.BoolVar(&testCigarCreation, "test", false, "Boolean to run basic tests")
	flag.StringVar(&server, "server", "localhost", "Server address")
	flag.StringVar(&dbPath, "db-path", "/cigarland/db", "Path for database")
	flag.StringVar(&dbName, "db-name", "cigar.db", "Name of the database file")
	flag.StringVar(&apiPrefix, "api-prefix", "api", "Name of api prefix for all api paths")
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.Parse()
}
