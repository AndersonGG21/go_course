package function

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Add(a, b int) int {
	return a + b
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Printf("%s - ", value)
	}
}

func Calc(op Operation, x, y float64) (float64, error) {

	switch op {
	case 0:
		return (x + y), nil
	case 1:
		return (x - y), nil
	case 2:
		if y == 0 {
			return 0, errors.New("you can't divide by zero")
		}
		return (x / y), nil
	case 3:
		return x * y, nil
	default:
		return 0, errors.New("you selected a wrong option")
	}
}

func Split(a, b int) (x, y int) {
	x = a / b
	y = a % b
	return
}

func MSum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
