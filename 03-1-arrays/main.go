package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var myArray [6]int
	stringArray := [6]string{"a", "b", "c", "d", "e", "f"}

	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
	myArray[4] = 5
	myArray[5] = 6

	fmt.Printf("The length of the integer arrays is: %d\n", len(myArray))
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myArray, unsafe.Sizeof(myArray), unsafe.Sizeof(myArray)*8)
	fmt.Println(myArray)
	fmt.Println(stringArray)
}
