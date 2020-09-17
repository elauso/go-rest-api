package service

import (
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/model"
)

type ProductService struct {
	productDao *model.ProductDao
}

func NewProductService(pd *model.ProductDao) *ProductService {
	return &ProductService{pd}
}

func (ps *ProductService) List() ([]*model.Product, error) {
	return ps.productDao.List()
}

func (ps *ProductService) Create(p *model.Product) (*model.Product, error) {
	return ps.productDao.Insert(p)
}
