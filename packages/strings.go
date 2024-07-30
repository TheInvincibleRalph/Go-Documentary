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

		//returns the number of bytes read
		size := r.Size()
		fmt.Println(size)

		//returns the number of bytes of the unread portion of the Reader
		len := r.Len()
		fmt.Println(len)
	}

	//starts reading from the 10th byte (which is the offset parameter supplied)
	n, _ := r.ReadAt(buffer, 10)
	fmt.Printf("ReadAt: %s\n", buffer[:n])

	//reads a single byte from the Reader
	y, _ := r.ReadByte()
	fmt.Printf("ReadByte: %c\n", y)

	// Seek 7 bytes from the current position
	q, _ := r.Seek(7, io.SeekCurrent)
	fmt.Println("Position after seeking 7 bytes forward:", q) // returns 21

	// Seek 2 bytes from the start
	x, _ := r.Seek(2, io.SeekStart)
	fmt.Printf("Seek: %d\n", x) //returns 2

	// Seek 1 byte before the end
	p, _ := r.Seek(-1, io.SeekEnd)
	fmt.Println("Position after seeking to 1 byte before end:", p) // Output: 13

}
