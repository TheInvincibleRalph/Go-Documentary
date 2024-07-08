package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	// Set a deadline for the context 1500 milliseconds from now
	deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)
	defer cancelCtx()

	printCh := make(chan int)
	// Start a goroutine to process numbers from the channel
	go doAnother(ctx, printCh)

	// Loop to send numbers 1, 2, and 3 into the printCh channel
	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			// Successfully sent num to printCh, now sleep for 1 second
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			// Context's deadline or cancellation is reached, exit the loop
			break
		}
	}

	// Cancel the context to clean up resources
	cancelCtx()

	// Wait for a short moment to allow the doAnother goroutine to finish
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Printf("doAnother: finished\n")
			return
		case num := <-printCh:
			fmt.Printf("doAnother: %d\n", num)
		}
	}
}

func main() {
	ctx := context.Background()
	doSomething(ctx)
	fmt.Println("main: finished")
}
