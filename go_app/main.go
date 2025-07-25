package main

import (
	"flag"
	"fmt"
	"log"
	//"net/url"
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
		router := gin.Default()
		SetRoutesAndRun(router)
	}
}

func ParseArgs() {
	flag.BoolVar(&createDB, "create-db", false, "Boolean toggle to create a new DB")
	flag.BoolVar(&startServer, "start-server", false, "Boolean toggle to start the server")
	flag.BoolVar(&testCigarCreation, "test", false, "Boolean to run basic tests")
	flag.StringVar(&server, "server", "localhost", "Server address")
	flag.StringVar(&dbPath, "db-path", "/cigarland/db", "Path for database")
	flag.StringVar(&dbName, "db-name", "cigar.db", "Name of the database file")
	flag.StringVar(&apiPrefix, "api-prefix", "", "Name of api prefix for all api paths")
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.Parse()
}
