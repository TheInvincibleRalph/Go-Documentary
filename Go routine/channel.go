package main

import (
	"fmt"
	"sync"
)

func generateNumbers(total int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("sending %d to channel\n", idx)
		ch <- idx
	}
}

func printNumbers(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("read %d from channel\n", num)
	}
}

func main() {
	var wg sync.WaitGroup
	numberChan := make(chan int)

	wg.Add(2)
	go printNumbers(numberChan, &wg)
	generateNumbers(3, numberChan, &wg) //generateNumbers is not a goroutine in order to prevent the program from proceeding to close() before any of the goroutine are done.
	close(numberChan)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
