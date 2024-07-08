package main

import "fmt"

func main() {

	channel := make(chan string, 3) //buffered channel

	characters := []string{"a", "b", "c"}

	for _, s := range characters {
		select {
		case channel <- s:
		}
	}

	close(channel)

	for result := range channel {
		fmt.Println(result)
	}
}
