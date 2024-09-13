package main

import "fmt"

func main() {
	age := 23

	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are not an adult")
	}

	// You can also use the else if statement
	fmt.Println("\nElse if statement")
	if age == 18 {
		fmt.Println("You are legal xd")
	} else if age >= 23 {
		fmt.Println("You are a semi-adult")
	} else {
		fmt.Println("You are a child")
	}

	// You can also use the short statement
	fmt.Println("\nShort statement")

	if value := false; !value {
		fmt.Println("This will not print")
	} else {
		fmt.Println("This will print")
	}

	// You can also use the switch statement
	fmt.Println("\nSwitch statement")

	switch age {
	case 18:
		fmt.Println("You are legal xd")
	case 23:
		fmt.Println("You are a semi-adult")
	default:
		fmt.Println("You are a child")
	}

	// You can also use the switch statement without a condition
	fmt.Println("\nSwitch without condition")
	switch {
	case age == 18:
		fmt.Println("You are legal xd")
	case age >= 23:
		fmt.Println("You are a semi-adult")
	default:
		fmt.Println("You are a child")
	}
}
