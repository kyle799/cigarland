package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
		db := OpenDB(dbName)
		InitializeDBTables(db, tables)

	}
	if startServer {
		log.Printf("starting webserver at %s:%d", server, port)
		router := gin.Default()
		router.GET("/test", test)
		router.Run(fmt.Sprintf("%s:%d", server, port))
	}
}

func ParseArgs() {
	flag.BoolVar(&createDB, "create-db", true, "Boolean toggle to create a new DB")
	flag.BoolVar(&startServer, "start-server", false, "Boolean toggle to start the server")
	flag.StringVar(&server, "server", "localhost", "Server address")
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.StringVar(&dbPath, "db-path", "/cigarland/db", "Path for database")
	flag.StringVar(&dbName, "db-name", "cigar.db", "Name of the database file")
	flag.Parse()
}
