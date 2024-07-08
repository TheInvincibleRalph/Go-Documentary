package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//CASE STUDY: Fetching Data from a Remote API

/*
Scenario

You are building a Go program that fetches data from a remote API.
The program needs to handle various types of errors, such as network issues, invalid responses, and rate limiting.
Each error should be handled differently.

Implementation

Define Custom Error Types:
We'll define custom error types to handle specific error scenarios.

Implement the Request Function:
This function will simulate making a network request and potentially returning different types of errors.

Error Handling in main:
The main function will handle these errors based on their type and properties.

This code provides a robust error-handling mechanism, ensuring different types of errors are appropriately managed and communicated to the user.
*/

// Custom error type for network errors
type NetworkError struct {
	msg       string
	temporary bool
}

func (e *NetworkError) Error() string {
	return e.msg
}

func (e *NetworkError) Temporary() bool {
	return e.temporary
}

// Custom error type for invalid responses
type InvalidResponseError struct {
	msg string
}

func (e *InvalidResponseError) Error() string {
	return e.msg
}

// Custom error type for rate limiting
type RateLimitError struct {
	msg        string
	retryAfter time.Duration
}

func (e *RateLimitError) Error() string {
	return e.msg
}

func (e *RateLimitError) RetryAfter() time.Duration {
	return e.retryAfter
}

// Simulated request function

/*
This function simulates fetching data and randomly returns different types of errors:
rand.Intn(4) generates a random number between 0 and 3.
Depending on the random number, it either returns nil (no error) or a specific error type:
0: No error.
1: A temporary NetworkError.
2: An InvalidResponseError.
3: A RateLimitError with a 2-second retry duration.
The default case returns a generic error.
*/

func fetchData() error {
	// Simulating different error scenarios
	switch rand.Intn(4) {
	case 0:
		return nil
	case 1:
		return &NetworkError{msg: "Network timeout", temporary: true}
	case 2:
		return &InvalidResponseError{msg: "Invalid response format"}
	case 3:
		return &RateLimitError{msg: "Rate limit exceeded", retryAfter: 2 * time.Second}
	default:
		return errors.New("unknown error")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	err := fetchData() //This line calls fetchData and stores the returned error (if any) in the variable err.
	if err != nil {
		fmt.Println("Error:", err)

		//TYPE SWITCH
		switch e := err.(type) {
		case *NetworkError:
			if e.Temporary() {
				fmt.Println("This is a temporary network error. Please try again.")
			} else {
				fmt.Println("This is a permanent network error. Please check your connection.")
			}
		case *InvalidResponseError:
			fmt.Println("The response from the server was invalid. Please contact support.")
		case *RateLimitError:
			fmt.Printf("Rate limit exceeded. Retry after %v seconds.\n", e.RetryAfter().Seconds())
			time.Sleep(e.RetryAfter())
			// Retry fetching data
			err = fetchData()
			if err != nil {
				fmt.Println("Error on retry:", err)
				os.Exit(1)
			}
			fmt.Println("Successfully fetched data on retry.")
		default:
			fmt.Println("An unknown error occurred.")
		}
		os.Exit(1)
	}

	fmt.Println("Successfully fetched data.")
}

/*
TYPE ASSERTION & TYPE SWITCH

Type assertion in Go is a mechanism used to extract the underlying concrete value of a certain type from an interface.
Interfaces in Go can hold values of any type. When we want to access a specific type from an interface, we use type assertion.

A type switch is a specialized form of switch statement in Go that allows us to test types of interface values.
It is particularly useful when we need to perform different actions based on the type of an interface value.

EXAMPLE

for _, shape := range shapes {
		switch s := shape.(type) {
		case Circle:
			fmt.Printf("Circle with radius %.2f has area: %.2f\n", s.Radius, s.Area())
		case Rectangle:
			fmt.Printf("Rectangsle with width %.2f and height %.2f has area: %.2f\n", s.Width, s.Height, s.Area())
		default:
			fmt.Println("Unknown shape")

*/
