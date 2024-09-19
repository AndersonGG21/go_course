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
			return 0, errors.New("You can't divide by zero")
		}
		return (x / y), nil
	case 3:
		return x * y, nil
	default:
		return 0, errors.New("You selected a wrong option")
	}

}
