package product

import "nafiz/domain"

type service struct {
	prdctReppo ProductRepo
}

func NewService(prdctRepo ProductRepo) Service {

	return &service{
		prdctReppo: prdctRepo,
	}
}

func (svc *service) Create(prdct domain.Product) (*domain.Product, error) {
	return svc.prdctReppo.Create(prdct)

}

func (svc *service) Get(id int) (*domain.Product, error) {
	return svc.prdctReppo.Get(id)
}
func (svc *service) List(page, limit int64) ([]*domain.Product, error) {

	return svc.prdctReppo.List(page, limit)

}
func (svc *service) Update(prdct domain.Product) (*domain.Product, error) {

	return svc.prdctReppo.Update(prdct)
}
func (svc *service) Delete(id int) error {

	return svc.prdctReppo.Delete(id)
}

func (svc *service) Count() (int64, error) {
	return svc.prdctReppo.Count()
}
