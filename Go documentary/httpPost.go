package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	payload := map[string]string{"message": "Hello, World!"} //craetes a map
	jsonData, err := json.Marshal(payload)                   //marshals the map into a json format (a byte slice type)

	//creates the request
	req, err := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonData)) //creates a post request to the url provided with jsonData as the request body. The body is wrapped in a bytes.NewBuffer to convert the json byte slice into a readable stream
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json") //sets the Content-Type header of the request to application/json, indicating that the request body contains JSON data.

	//sends the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	//reads the response body into a byte slice and prints it as a string
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
