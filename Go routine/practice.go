package main

import (
	"fmt"
	// "time"
)

func main() {
	channel := make(chan string)

	go func() {
		channel <- "Hi, good morning!"
	}()

	receiver := <-channel
	fmt.Println(receiver)

	// time.Sleep(time.Second * 2)

	sayHello()
}

func sayHello() {
	fmt.Println("Hi!")
}
