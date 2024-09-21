package main

import (
	"encoding/json"
	"fmt"
	"go_course/06-1-struct/structs"
)

func main() {
	var product structs.Product

	product.ID = 1
	product.Name = "Samsung S9 F3"

	fmt.Println(product)

	product2 := structs.Product{}
	product2.Name = "Xiaomi Redmi K20"

	fmt.Println(product2)

	product3 := structs.Product{
		ID:   1,
		Name: "Iphone 16 Pro Max",
		Type: structs.Type{
			ID:          1,
			Code:        "N/A",
			Description: "SmarthPhones",
		},
		Price: 5000,
		Count: 3,
		Tags:  []string{"Apple", "Expensive", "Money"},
	}

	product4 := structs.Product{
		ID:   2,
		Name: "Iphone 15 Pro Max",
		Type: structs.Type{
			ID:          1,
			Code:        "N/A",
			Description: "SmarthPhones",
		},
		Price: 2000,
		Count: 2,
		Tags:  []string{"Apple", "Expensive", "Money"},
	}

	json, err := json.Marshal(product3)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("JSON", string(json))

	fmt.Println(product3)

	fmt.Println("Product 3 Total Price", TotalPrice(product3))
	fmt.Println("Product 3 Total Price", product3.TotalPrice())
	product3.AddTags("Tag1", "Tag2", "Tag3")
	fmt.Println("Product 3 New Tags", product3)

	//I can also create a 'constructor method for Cart Struct'
	cart := structs.NewCart(1)

	cart.AddProducts(product3, product4)
	fmt.Println("\nCart", cart)
	fmt.Println("Order Total", cart.OrderTotalPrice())

}

func TotalPrice(p structs.Product) float64 {
	return p.Price * float64(p.Count)
}
