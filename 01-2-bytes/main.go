package main

import (
	"fmt"
)

func main() {

	var A byte = 'A'
	var a byte = 'a'
	var R byte = 82
	var r byte = 114
	vector := []byte{A, 110, 100, 101, 114, 115, 111, 110}

	fmt.Printf("A-Type: %T, A-ASCII: %d, a-type: %T, a-ASCII: %d", A, A, a, a)
	fmt.Printf("\nR-Type: %T, R-ASCII: %d, r-type: %T, r-ASCII: %d", R, R, r, r)

	//To convert a byte to a string, we can use the string() function
	fmt.Printf("\nA-String: %s, a-String: %s, R-String: %s, r-String: %s", string(A), string(a), string(R), string(r))

	//To convert a byte slice to a string, we can use the string() function
	fmt.Printf("\nVector-String: %s\n", string(vector))
}
