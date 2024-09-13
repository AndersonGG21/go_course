package main

import "fmt"

//Implement a program in go to enter values by console until zero is added and the program ends.

func main() {

	var mySlice []int
	var input int = 1

	for input != 0 {
		fmt.Println("Ingrese un valor: ")
		fmt.Scan(&input)
		mySlice = append(mySlice, input)
	}
	mySlice = mySlice[:len(mySlice)-1] //I know it's not the most optimal way, I just wanted to review
	fmt.Println("Valores ingresados:", mySlice)

	mySlice = mySlice[:0]

	//Using 'while'
	for {
		fmt.Println("Ingrese un valor: ")
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		mySlice = append(mySlice, input)
	}

	fmt.Println("Valores ingresados:", mySlice)
}
