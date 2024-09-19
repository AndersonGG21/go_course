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
}
