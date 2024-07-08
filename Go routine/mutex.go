package main

import (
	"fmt"
	"sync"
)

var CommonParam = 0 //declares a global variable and initialixe it to 0. This variable will be shared and incremented by by multiple goroutines

func Increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()     //this lines locks the shared variable ensuring that the code block is executed by only one goroutine at a time.
	CommonParam += 1 //one goroutine handles this at a time after which the shared resource is unlocked by the line of code below.
	mutex.Unlock()
	wg.Done() //the particular goroutine declares done and decrements WaitGroup counter by 1
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 100; i++ {
		w.Add(1)
		go Increment(&w, &m) //this line of code passes in the memory location of the objects of the Increment function thereby spawning a new go routine.
	}

	w.Wait() //this line blocks the main goroutine until the WaitGroup counter goes back to 0. This means that the main goroutine will wait here until all 100 goroutines have called wg.Done()
	fmt.Println("CommonParam: ", CommonParam)
}
