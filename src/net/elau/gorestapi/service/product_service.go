package service

import (
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/model"
)

type ProductService struct {
	productDao *model.ProductDao
}

func NewProductService(productDao *model.ProductDao) *ProductService {
	return &ProductService{productDao}
}

func (ps *ProductService) List() ([]*model.Product, error) {
	return ps.productDao.List()
}
