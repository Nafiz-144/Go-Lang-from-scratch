package database

// ------------------ PRODUCT STRUCT ------------------
// এটা হচ্ছে data structure (model)
// JSON ↔ Go struct mapping হয়

type Product struct {
	ID          int     `json:"id"`          // product id
	Title       string  `json:"title"`       // product name
	Description string  `json:"description"` // details
	Price       float64 `json:"price"`       // price
	ImgUrl      string  `json:"imgUrl"`      // image URL
}

// ------------------ IN-MEMORY DATABASE ------------------
// slice ব্যবহার করে temporary database বানানো হয়েছে
// ⚠️ server restart হলে data চলে যাবে

var ProductList []Product

// ------------------ INIT FUNCTION ------------------
// init() automatically run হয় main() এর আগে
// এখানে initial data insert করা হচ্ছে

func init() {

	Pr1 := Product{
		ID:          1,
		Title:       "Mango",
		Description: "I love Mango, That's an interesting fruit",
		Price:       50,
		ImgUrl:      "https://afrisunorchards.com/wp-content/uploads/2023/12/mangoes-495x400.jpg",
	}

	// slice এ add করা হচ্ছে
	ProductList = append(ProductList, Pr1)
}
