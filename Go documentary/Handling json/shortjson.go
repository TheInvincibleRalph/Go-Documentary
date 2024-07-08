package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"` //json tag is used to determine what fields in the json maps to what fields in the structure and vice versa.
	Completed bool   `json:"completed"`
}

func main() {

	url := "http://jsonplaceholder.typicode.com/todos/1/"

	response, err := http.Get(url) //response is pointer to an http.Response struct that contains the HTTP response from the server.
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK { //this checks if the HTTP status code of the response is 200 OK, which means the request was successful. It then goes ahead to read the body of the page.

		todoItem := Todo{}                               //this creates an instance of the Todo struct.
		json.NewDecoder(response.Body).Decode(&todoItem) //NewDecoder returns a new decoder that reads from the parameters we pass into it.
		fmt.Printf("Data from API: %+v", todoItem)       //the &todoItem argument is a pointer to the todoItem struct, allowing json.Unmarshal to modify its fields directly.
	}
	/*
	   In Go, the %+v verb in the fmt.Printf function is used
	   to print Go structs in a more detailed and readable format.
	   When you use %+v, the struct's field names along with their
	   values are included in the output.
	*/
}
