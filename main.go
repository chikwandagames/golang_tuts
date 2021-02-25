package main

import (
	"fmt"
	"runtime"
)

// byte is and alias for uint8
// rune is and alias for uint32

/*
	Package runtime contains operations that interact with Go's runtime system,
	such as functions to control goroutines. It also includes the low-level type
	information used by the reflect package
*/
func main() {
	// Runtime package give details about the
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
