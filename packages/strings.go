package main

import (
	"fmt"
	"io"
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

	//READER TYPE

	r := strings.NewReader("Hello, Reader!") //initializes a Reader instance

	buffer := make([]byte, 8)

	//reads from the Reader in chunks of 8 bytes
	for {
		n, err := r.Read(buffer) //Read reads data from the Reader instance into the buffer
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			break
		}
		fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
	}
}
