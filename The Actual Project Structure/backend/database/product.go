package database

// Product structure (Model)
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

var productList []Product

func Store(p Product) Product {
	p.ID = len(productList) + 1

	productList = append(productList, p)
	return p

}
func List() []Product {

	return productList
}
func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}
func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product

		}
	}

}

func Delete(productID int) {
	var tempList []Product = make([]Product, 0)

	for _, p := range productList {
		if p.ID != productID {
			tempList = append(tempList, p)

		}
	}
	productList = tempList

}

func init() {
	Pr1 := Product{
		ID:          1,
		Title:       "Mango",
		Description: "I love Mango, That's an interesting fruit",
		Price:       50,
		ImgUrl:      "https://afrisunorchards.com/wp-content/uploads/2023/12/mangoes-495x400.jpg",
	}
	productList = append(productList, Pr1)
}
