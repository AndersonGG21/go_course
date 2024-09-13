package main

import "fmt"

func main() {

	myMap := map[int]string{
		10: "notebook",
		15: "tv",
		21: "heladera",
		27: "monitor",
		35: "camara",
	}

	var resultSlice []string
	input := 0

	for {
		fmt.Println("Insert a key:")
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		if _, exits := myMap[input]; exits {
			resultSlice = append(resultSlice, myMap[input])
		} else {
			resultSlice = append(resultSlice, "Not found")
		}
	}

	fmt.Println(resultSlice)

}
