package database

// Product structure (Model)
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

// Global slice to store all products (in-memory database)
var ProductList []Product

// ====================== INIT FUNCTION ======================

// Runs automatically before main()
// Used here to insert initial data
func init() {
	Pr1 := Product{
		ID:          1,
		Title:       "Mango",
		Description: "I love Mango, That's an interesting fruit",
		Price:       50,
		ImgUrl:      "https://afrisunorchards.com/wp-content/uploads/2023/12/mangoes-495x400.jpg",
	}
	ProductList = append(ProductList, Pr1)
}
