package main

import (
	"fmt"
)

func main() {
	// A string is a sequence of bytes (slice of bytes)
	// Strings are immutable
	s := "hello golang"
	fmt.Println(s)
	fmt.Printf("%T \n", s)

	bs := []byte(s)
	fmt.Println("ASCII values of the string:", bs)
	fmt.Printf("%T \n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("ASCII: %v, utf8 or rune: %#U \n", s[i], s[i])
	}

	for i, v := range s {
		fmt.Printf("index: %d, hev: %#x, ASCII: %v \n", i, v, s[i])
	}
}
