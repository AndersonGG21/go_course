package main

import (
	"fmt"
	"go_course/05-1-functions/function"
)

func main() {
	var a int = 9
	b := 2

	fmt.Println(function.Add(a, b))
	function.RepeatString(3, "I'm the best 🐐")
	res, err := function.Calc(function.DIV, 2, 0)

	fmt.Println()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("The result of the operation is", res)
	}

	fmt.Println(function.MSum(1, 2, 3, 4, 5))

	fmt.Println("\nFactory functions:")

	sumFunc := function.FactoryOperation(function.SUM)
	fmt.Println(sumFunc(1, 2))

	subFunc := function.FactoryOperation(function.SUB)
	fmt.Println(subFunc(1, 2))

	mulFunc := function.FactoryOperation(function.MUL)
	fmt.Println(mulFunc(1, 2))

	divFunc := function.FactoryOperation(function.DIV)
	fmt.Println(divFunc(4, 2))

}
