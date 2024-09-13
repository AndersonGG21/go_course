package main

//A slice is a segment of an array. It is a flexible-length array that allows elements to be added or removed
//What I know as dynamic array
import (
	"fmt"
)

func main() {
	var mySlice []int
	var myArray [6]int

	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
	myArray[4] = 5
	myArray[5] = 6

	fmt.Printf("length: %d, value: %v\n", len(mySlice), mySlice)

	mySlice = append(mySlice, 1, 2, 3, 4, 5)
	fmt.Printf("length: %d, value: %v\n", len(mySlice), mySlice)

	mySlice2 := myArray[:3]
	fmt.Printf("length: %d, value: %v\n", len(mySlice2), mySlice2)

	fmt.Println("\n Memory address")
	fmt.Printf("array: %p\n", &myArray[0])
	fmt.Printf("mySlice2: %p\n", &mySlice2[0])

	mySlice2[0] = 100
	fmt.Println(myArray[0])

	//Declaring a slice using the make function

	mySlice3 := make([]int, 3, 5)
	fmt.Printf("length: %d, capacity: %d, value: %v\n", len(mySlice3), cap(mySlice3), mySlice3)
	//The capacity is the maximum number of elements that the slice can hold without having to allocate more memory

	mySlice3 = append(mySlice3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("length: %d, capacity: %d, value: %v\n", len(mySlice3), cap(mySlice3), mySlice3)
}
