package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//creates the request
	client := &http.Client{}                                           //this creates a request object, a struct type.
	req, err := http.NewRequest("GET", "https://www.google.com/", nil) //(METHOD, URL, BODY(usually nil for get but an io.Reader for POST and PUT))
	if err != nil {
		log.Fatal(err)

	}

	//modifies the query
	q := req.URL.Query()          //extracts the query parameters from the URL
	q.Add("sku", "How are you?")  //adds a query parameter 'sku' with the value 'ABC123'
	req.URL.RawQuery = q.Encode() //encodes the query parameters back into the URL

	//sends the request
	response, err := client.Do(req) //sends the http request and receives the response
	if err != nil {
		fmt.Println(err)
		defer response.Body.Close()
	}

	//read the response and print as a string
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

}
