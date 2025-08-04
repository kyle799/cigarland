package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MakeDBHandler(db *gorm.DB, handlerFunc func(*gorm.DB, *gin.Context)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		handlerFunc(db, ctx)
	}
}

func SelectWhereHandler(db *gorm.DB, ctx *gin.Context) {
	log.Printf("Received request to query the cigar table with a filter")
	bodyContent, readErr := io.ReadAll(ctx.Request.Body)
	if readErr != nil {
		responseMap := map[string]string{
			"message": fmt.Sprintf("Error reading request body: %s", readErr),
		}
		response, marshalErr := json.Marshal(responseMap)
		if marshalErr != nil {
			ctx.Status(500)
			fmt.Fprintf(ctx.Writer, "Failed to read request body: %s\nFailed to marshal json response :%s", readErr, marshalErr)
			return
		}
		ctx.Data(500, "application/json", response)
		return
	}
	log.Printf("Request Body:\n%s", bodyContent)
	query := QueryPayload{}
	unmarshalErr := json.Unmarshal(bodyContent, &query)
	if unmarshalErr != nil {
		responseMap := map[string]string{
			"message": fmt.Sprintf("Failed to unmarshal payload: %s", unmarshalErr),
		}
		response, _ := json.Marshal(responseMap)
		ctx.Data(500, "application/json", response)
		return
	}
	log.Printf("Querying table: %s with: %+v", query.Table, query.Filters)
	queryResult, _ := QueryDB(db, query.Table, query.Filters...)
	responseBody, err := json.Marshal(queryResult)
	if err != nil {
		ctx.String(500, "application/text", fmt.Sprintf("Error generating query response: %s", err))
		return
	}
	ctx.Data(200, "application/json", responseBody)
}

func HandleCreateCigar(ctx *gin.Context, db *gorm.DB) {
	log.Printf("receiving request to insert cigar")
	log.Printf("making byte slice from requst content length: %d", ctx.Request.ContentLength)
	// bodyContent := make([]byte, ctx.Request.ContentLength)
	log.Printf("reading request body into byte slice")
	bodyContent, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Printf("Error reading from request body: %s", err)
	}
	// log.Printf("Read bytes from requeset body: %d", readBytes)
	log.Printf("Request Body: %s", bodyContent)
	log.Printf("creating cigar payload struct")
	cigarCreatePayload := CigarCreatePayload{}
	log.Printf("unmarshaling payload into cigarpayload struct")
	err = json.Unmarshal(bodyContent, &cigarCreatePayload)
	if err != nil {
		log.Printf("error unmarshaling json into cigarCreatePayload struct: %s", err)
	} else {
		log.Printf("CigarCreatePayload: %+v", cigarCreatePayload)
	}
	cigarInfo := cigarCreatePayload.Cigars
	returnCode, jsonResponse := CreateCigars(db, cigarInfo)
	ctx.JSON(returnCode, jsonResponse)
}

func CreateCigars(db *gorm.DB, cigars []*Cigar) (returnCode int, jsonBody map[string]any) {
	jsonBody = make(map[string]any)
	cigarCount := len(cigars)
	if cigarCount > 0 {
		log.Printf("Creating cigars. Count: %d", cigarCount)
		dbSession := db.Session(&gorm.Session{CreateBatchSize: cigarCount})
		result := dbSession.Create(&cigars)
		if result.Error != nil {
			jsonBody["error"] = result.Error
			jsonBody["rows_effected"] = result.RowsAffected
			log.Printf("Insertion error occured: %s", result.Error)
			// ctx.JSON(500, jsonResponse)
			returnCode = 500
		} else {
			jsonBody["rows_effected"] = result.RowsAffected
			log.Printf("Updated Cigar Table, Rows Affected: %d", result.RowsAffected)
			returnCode = 200
			// ctx.JSON(200, jsonResponse)
		}
	} else {
		log.Printf("Cigar creation called with no cigars provided in the request")
		jsonBody["message"] = "No cigars provided to create/update"
		returnCode = 200
	}
	return returnCode, jsonBody
}

/*
 */
func HandleQueryCigarTable(db *gorm.DB, ctx *gin.Context) {
	cigars := QueryCigarTable(db)
	ctx.JSON(200, cigars)
}

func QueryCigarTable(db *gorm.DB) []*Cigar {
	log.Printf("Received request to query the cigar table")
	cigars := []*Cigar{}
	db.Find(&cigars)
	log.Printf("Number of cigars found: %d", len(cigars))
	return cigars
}
