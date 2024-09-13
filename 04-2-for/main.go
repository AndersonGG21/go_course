package main

import "fmt"

// For arrays the first value is the index and the second the value, if I declare only one variable it takes the value of the index.
func main() {

	sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
	}

	fmt.Printf("La suma es %d\n", sum)

	for sum < 1000 {
		sum += sum
	}

	fmt.Printf("La suma es %d\n", sum)

	for {
		if sum > 2000 {
			break
		}
		sum++
	}

	fmt.Printf("La suma es %d\n", sum)

	//To irate an array
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := range arr {
		fmt.Printf("Value in index %d is %d\n", i, arr[i])
	}

	myMap := map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
	}
	fmt.Println()

	//To iterate a map
	for key, value := range myMap {
		fmt.Printf("Value with key %s is %d\n", key, value)
	}

	myMap2 := map[string][]int{
		"arr1": {1, 2, 3},
		"arr2": {4, 5, 6},
	}

	fmt.Println("\nMap with arrays as value:")
	for _, value := range myMap2 {
		for _, n := range value {
			fmt.Println(n)
		}
	}
}
