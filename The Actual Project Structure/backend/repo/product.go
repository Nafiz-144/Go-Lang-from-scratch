package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)

	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

// constructor or constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProduct(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1

	r.productList = append(r.productList, &p)
	return &p, nil
}
func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}
	return nil, nil
}
func (r *productRepo) List() ([]*Product, error) {

	return r.productList, nil
}
func (r *productRepo) Delete(productID int) error {
	var tempList []*Product

	for _, p := range r.productList {
		if p.ID != productID {
			tempList = append(tempList, p)

		}
	}
	r.productList = tempList
	return nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product

		}
	}
	return &product, nil
}
func generateInitialProduct(r *productRepo) {
	Pr1 := &Product{
		ID:          1,
		Title:       "Mango",
		Description: "I love Mango, That's an interesting fruit",
		Price:       50,
		ImgUrl:      "https://afrisunorchards.com/wp-content/uploads/2023/12/mangoes-495x400.jpg",
	}
	r.productList = append(r.productList, Pr1)
}
