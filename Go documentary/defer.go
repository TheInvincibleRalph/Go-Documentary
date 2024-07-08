// The defer keyword  allows us to postpone the execution of a statement or a function until the end of the enclosing (calling) function
//It executes something (a function or an expression) when the enclosing function returns (after every returns and even when an error
//occurred in the midst of executiong the function, not only return at the end of the function, but before the })
//Why after? Because the return statement itself can be an expression which does something instead of only giving back 1 or more variable.

/*
In Go, defer is a keyword used to schedule a function call to be executed when the surrounding function (the one containing the defer statement) exits
*/
package main

import "fmt"

func main() {
	Function1()

	// Multiple defer statements in a function are stacked in a Last In, First Out (LIFO) manner
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")

}

func Function1() {
	fmt.Printf("In Function1 at the top\n")
	defer Function2()
	fmt.Printf("In Function2 at the bottom!\n")
}

func Function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!\n")
}
