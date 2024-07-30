package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	//append multiple strings
	builder.WriteString("Hello, ")
	builder.WriteString("World!")

	//using WriteRune to add a single rune character
	builder.WriteRune(' ')
	builder.WriteString("I love Go.")

	res := builder.String()
	fmt.Println(res)

	//resets the builder to reuse it
	builder.Reset()
	builder.WriteString("Buffer already resetted")
	fmt.Println(builder.String()) //Output: Buffer already resetted.

	s := builder.Cap()
	fmt.Println(s)

	builder.Grow(20)

}
