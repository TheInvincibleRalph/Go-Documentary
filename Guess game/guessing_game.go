package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//Create secret number
	secret := getRandomNumber()

	for matching := false; !matching; {

		//Get user input
		guess := getUserInput()

		//Make comparison (secret vs guess)
		matching = isMatching(secret, guess)
	}
}

// Create a function to make comparison between the secret and guess number
func isMatching(secret, guess int) bool {
	if guess > secret {
		fmt.Println("Your guess is too big")
		return false
	} else if guess < secret {
		fmt.Println("Your guess is too small")
		return false
	} else {
		fmt.Println("YOU GOT IT!")
		return true
	}
}

// Create a function to get user input
func getUserInput() int {
	fmt.Print("Please print your guess: ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Failed to parse your input")
	} else {
		fmt.Println("You guessed:", input)
	}
	return input
}

// Create a function to generate the random number
func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}

// func grade() {
// 	score := 120

// 	if score < 49 {
// 		fmt.Println("Your grade  is D")
// 	} else if score > 49 && score <= 59 {
// 		fmt.Println("Your grade  is C")
// 	} else if score > 59 && score <= 69 {
// 		fmt.Println("Your grade  is B")
// 	} else if score > 69 && score <= 100 {
// 		fmt.Println("Your grade  is A")
// 	} else {
// 		fmt.Println("Score out of range")
// 	}
// }

// func main() {
// 	grade()
// }
