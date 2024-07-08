package main

import "fmt"

/*This is an array with undefined length (represented with "..."). The range keyword loops through the array
and returns the index of each item as an integer*/

func main() {
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
}
