package main

import (
	"fmt"
	"go_course/06-2-interfaces/vehicles"
)

func main() {
	fmt.Println("Vehicles")

	car := vehicles.Car{
		Time: 120,
	}

	fmt.Println(car.Distance())

	vArray := []string{vehicles.CarV, "MOTORCYCLE", "TRUCK", "MOTORCYCLE", "TRUCK", "PLANE"}

	d := 400
	for _, v := range vArray {
		fmt.Printf("Vehicle: %s\n", v)
		vehicle, err := vehicles.New(v, d)
		if err != nil {
			fmt.Print(err.Error())
			continue
		}

		fmt.Printf("Distance traveled by %s is %.2f\n\n", v, vehicle.Distance())

	}
}
