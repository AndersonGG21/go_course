package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("Integer:", myIntVar)

	var myUnsignedIntVar uint
	myUnsignedIntVar = 69
	fmt.Println("Unsigned Integer:", myUnsignedIntVar)

	var myStringVar string
	myStringVar = "Hello, World!"
	fmt.Println("String:", myStringVar)

	var varMyBoolVar bool
	varMyBoolVar = true
	fmt.Println("Boolean:", varMyBoolVar)

	fmt.Println("My string var address is:", &myStringVar)

	myIntVar2 := 42
	fmt.Println("Integer 2:", myIntVar2)

	fmt.Println()

	fmt.Println("****************")
	fmt.Println("***Constantes***")
	fmt.Println("****************")

	const myConstIntVar int = 42
	fmt.Println("Constant Integer:", myConstIntVar)

	const myConstStringVar string = "Hello, World!"
	fmt.Println("Constant String:", myConstStringVar)

	/*
		int8 - 8bits -128 a 127
		int16 - 16bits -32768 a 32767
		int32 - 32bits -2147483648 a 2147483647
		int64 - 64bits -9223372036854775808 a 9223372036854775807
	*/

	fmt.Println()

	fmt.Println("****************")
	fmt.Println("***Format*******")
	fmt.Println("****************")

	var myInt8Var int8
	fmt.Printf("Int8 default value: %d\n", myInt8Var)

	//Unsafe package is used to get the size of a variable
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myInt8Var, unsafe.Sizeof(myInt8Var), unsafe.Sizeof(myInt8Var)*8)

	fmt.Println()

	fmt.Println("****************")
	fmt.Println("***Conversions**")
	fmt.Println("****************")
	{
		floatVar := 3.14
		fmt.Printf("type: %T, value: %v\n", floatVar, floatVar)
		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar)

		intVar := 42
		fmt.Printf("type: %T, value: %v\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s\n", intStrVar, intStrVar)

		//Converting string to int
		strVar := "42"
		fmt.Printf("type: %T, value: %v\n", strVar, strVar)
		intVar2, _ := strconv.Atoi(strVar) //-> Another way is to use ParseInt adding the base, example: strconv.ParseInt(strVar, 10, 0)
		fmt.Printf("type: %T, value: %v\n", intVar2, intVar2)

	}
}
