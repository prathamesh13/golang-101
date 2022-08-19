// comments are allowed before package names
/*
used by tools that generate/read docs
*/

// This package does nothing. If if was, there would be more info here.
//
// possibly little more info here as well.

package main

import (
	"fmt"
)

// typical single line comment - similar to other languages
/*this will show warning in Visual Studio code and possibly other Go IDE since
each go package should have only one main function defined. we can still test
this sample code by running `go run comments.go`
*/
func main() {

	fmt.Println("comments in golang")

	var number = 42

	printValueAndAddress(number, &number)

}

/*
A multi line comments. This is used by golang compiler to generate documentation for
the go module. This will indicate the purpose of function, and describe function
parameters and return values
*/
func printValueAndAddress(value int, address *int) {

	fmt.Printf("Integer %v stored at %v", value, address)

}
