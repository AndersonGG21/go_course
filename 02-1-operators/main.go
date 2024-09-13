package main

import "fmt"

func main() {
	yearsOld := 32
	fmt.Println(yearsOld > 30)
	fmt.Println(yearsOld < 30)
	fmt.Println(yearsOld >= 32)
	fmt.Println(yearsOld <= 32)
	fmt.Println(yearsOld == 32)
	fmt.Println(yearsOld != 32)

	// Logical operators
	// OR

	fmt.Println("\n*** OR ***")
	fmt.Println(yearsOld > 30 || yearsOld < 30)
	fmt.Println(yearsOld > 30 || yearsOld == 32)
	fmt.Println(yearsOld < 30 || yearsOld == 32)
	fmt.Println(yearsOld < 30 || yearsOld > 32)

	// AND
	fmt.Println("\n*** AND ***")
	fmt.Println(yearsOld > 30 && yearsOld < 30)
	fmt.Println(yearsOld > 30 && yearsOld == 32)
	fmt.Println(yearsOld < 30 && yearsOld == 32)
	fmt.Println(yearsOld < 30 && yearsOld > 32)

	// NOT
	fmt.Println("\n*** NOT ***")
	fmt.Println(!(yearsOld > 30))
	fmt.Println(!(yearsOld < 30))
	fmt.Println(!(yearsOld == 32))
	fmt.Println(!(yearsOld != 32))

}
