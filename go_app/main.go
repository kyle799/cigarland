package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	ParseArgs()
	PopulateValueOperatorMap()
	dsn := buildDSN()
	if createDB {
		fmt.Printf("Starting creation of db\n")
		cigarDB = OpenDB(dsn)
		tables := CreateTableSchemas()
		InitializeDBTables(cigarDB, tables)
		SeedAdminUser(cigarDB)
	}
	if startServer {
		if cigarDB == nil {
			cigarDB = OpenDB(dsn)
		}
		InitOAuth()
		router := gin.Default()
		SetRoutesAndRun(router)
	}
}

func buildDSN() string {
	if dbDSN != "" {
		return dbDSN
	}
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "cigarland")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func ParseArgs() {
	flag.BoolVar(&createDB, "create-db", false, "Boolean toggle to create/migrate the DB schema")
	flag.BoolVar(&startServer, "start-server", false, "Boolean toggle to start the server")
	flag.BoolVar(&testCigarCreation, "test", false, "Boolean to run basic tests")
	flag.StringVar(&server, "server", "localhost", "Server address")
	flag.StringVar(&dbDSN, "db-dsn", "", "PostgreSQL DSN (overrides DB_HOST/DB_PORT/DB_USER/DB_PASSWORD/DB_NAME env vars)")
	flag.StringVar(&apiPrefix, "api-prefix", "", "Name of api prefix for all api paths")
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.Parse()
}
