package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	//"fmt"
	"net/http"
	"net/url"
	//"os"
)

//func TestCigarDBInsertion(dbName string) {
//	db := OpenDB(dbName)
//	cigars := CreateRandomCigars(10)
//	for _, cigar := range cigars {
//		db.Create(cigar)
//	}
//}

func CreateAndSendCigarCreationPayload(url *url.URL) {
	client := http.Client{}
	requestBody := &bytes.Buffer{}
	cigarPayload := CreateRandomCigarPayload()
	jsonContent, err := json.MarshalIndent(cigarPayload, "", "  ")
	if err != nil {
		lErr.Printf("Error marshaling payload to json: %s\n", err)
		return
	}
	requestBody.Write(jsonContent)
	lInfo.Printf("Content length of request: %d", requestBody.Len())
	lErr.Printf("%s", requestBody.Bytes())
	resp, err := client.Post(url.String(), "application/json", requestBody)
	if err != nil {
		lErr.Printf("Error received from the server: %s\n", err)
	}
	lInfo.Printf("Resopnse: %+v\n", resp)
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
	c := Cigar{}
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
	return &c
}

func QueryCigarTable(URL *url.URL) []*Cigar {
	cigars := []*Cigar{}
	client := http.Client{}
	lInfo.Printf("Querying url: %s", URL.String())
	response, _ := client.Get(URL.String())
	bodyContent, _ := io.ReadAll(response.Body)
	json.Unmarshal(bodyContent, &cigars)
	return cigars
}

func FilterCigarTableByRating(URL *url.URL) ([]*Cigar, error) {
	encoder := json.NewEncoder(lInfo.Writer())
	encoder.SetEscapeHTML(false)
	filter := SelectionFilter{
		Column:   "Rating",
		Value:    rand.Int() % 11,
		Operator: GetRandomStringKey(ValueOperatorMap[TypeInt]),
		// Operator: ">=",
		Logical: "",
	}
	if debug {
		encoder.Encode(filter)
	}
	payload := QueryPayload{Table: "Cigars", Filters: []SelectionFilter{filter}}
	if debug {
		encoder.Encode(payload)
	}
	payloadContent, _ := json.Marshal(&payload)
	payloadBuffer := &bytes.Buffer{}
	payloadBuffer.Write(payloadContent)
	client := http.Client{}
	response, err := client.Post(URL.String(), "application/json", payloadBuffer)
	if err != nil {
		return nil, fmt.Errorf("Error sending post to api: %w", err)
	}
	responseContent, _ := io.ReadAll(response.Body)
	queryResponse := []*Cigar{}
	_ = json.Unmarshal(responseContent, &queryResponse)
	return queryResponse, nil
}
