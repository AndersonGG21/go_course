package main

import "fmt"

//A map is an unordered collection of key-value pairs. Also known as an associative array, a hash table or a dictionary,
//maps are used to look up a value by its associated key.
//Each key must be unique and of the same type, and it can be of any type. The values can also be of any type.
//The zero value of a map is nil. A nil map has no keys, nor can keys be added.

func main() {

	myMap := make(map[int]string)

	myMap[1] = "One"
	myMap[2] = "Two"
	myMap[3] = "Three"

	fmt.Println(myMap)
	fmt.Println(myMap[2])

	//Deleting a key-value pair from a map
	delete(myMap, 2)
	fmt.Println(myMap)

	//Adding a key-value pair to a map
	myMap[2] = "Two"
	fmt.Println(myMap)

	//Checking if a key exists in a map
	value, exists := myMap[2]
	fmt.Println("Value:", value, "Exists:", exists)

	myMap2 := map[int]string{
		1: "Uno",
		2: "Dos",
		3: "Tres",
	}
	fmt.Println(myMap2)
}
