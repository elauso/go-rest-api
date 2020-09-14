package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/service"
)

type ProductRoute struct {
	productService *service.ProductService
}

func NewProductRoute(productService *service.ProductService) *ProductRoute {
	return &ProductRoute{productService}
}

func (pr *ProductRoute) List(w http.ResponseWriter, r *http.Request) {

	plist, err := pr.productService.List()
	if err != nil {
		log.Printf("Failed to retrieve products, %v", err)
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(plist)
	w.WriteHeader(200)
}
