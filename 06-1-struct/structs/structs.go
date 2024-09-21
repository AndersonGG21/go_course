package structs

type Product struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Type  Type    `json:"type"`
	Price float64 `json:"price"`
	Count int     `json:"count"`

	Tags []string `json:"tags"`
}

type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Cart struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"userID"`
	Products []Product `json:"products"`
}

func (product Product) TotalPrice() float64 {
	return product.Price * float64(product.Count)
}

func (product *Product) AddTags(tags ...string) {
	product.Tags = append(product.Tags, tags...)
}

func (car *Cart) AddProducts(products ...Product) {
	car.Products = append(car.Products, products...)
}

func (car Cart) OrderTotalPrice() float64 {
	var total float64
	for _, product := range car.Products {
		total += product.Price * float64(product.Count)
	}
	return total
}

func NewCart(userID uint) Cart {
	return Cart{UserID: userID}
}
