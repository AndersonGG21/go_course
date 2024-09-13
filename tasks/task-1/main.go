package main

import "fmt"

// Inside our code we are going to have 2 variables defined, which are going to be:
// license: if the person has a license
// age: the age of the person
//     If the person is over 15 years old and has a license, we must print a message that says “You can continue to advance”.
//     If the person is 15 years old or younger, or does not have a license, we must print a message that says “You can not continue circulating”.

func main() {

	license := false
	age := 15

	if age > 15 && license {
		fmt.Println("You can continue to advance")
	} else if age <= 15 || !license {
		fmt.Println("You can not continue circulating")
	}

}
