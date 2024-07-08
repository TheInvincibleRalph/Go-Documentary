package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//send an http request using the get method (just like initiating a connection)
	response, err := http.Get("https://www.google.com/teapot")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	//read the body of the response from the get request and print it out as a string
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
