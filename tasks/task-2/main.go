package main

import "fmt"

//Given an array array, we must run it and increment all its values by 20.

func main() {
	array := [5]int{4, 2, 5, 6, 7}

	for i := range array {
		array[i] = array[i] + 20
	}

	fmt.Println("Los valores del array son: ", array)

}
