package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*
	type error interface {
		Error() string
	}
*/

// Any type that implements Error() string is also of type error

func main() {
	// Create new file, returns a pointer to an error
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Returns a Reader
	r := strings.NewReader("Hello")

	// Copy reader "string" into file
	io.Copy(f, r)

	// Now read from the file
	// This wont work, because the read operation will run before the
	// write operation has completed
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
