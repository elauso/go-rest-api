package route

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/model"
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/service"
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/web/response"
	"github.com/gorilla/mux"
)

type ProductRoute struct {
	productService *service.ProductService
}

func NewProductRoute(ps *service.ProductService) *ProductRoute {
	return &ProductRoute{ps}
}

func (pr *ProductRoute) List(w http.ResponseWriter, r *http.Request) {

	plist, err := pr.productService.List()
	if err != nil {
		log.Printf("Failed to retrieve products, %v", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(plist)
}

func (pr *ProductRoute) Get(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var id uint64
	var err error

	if id, err = strconv.ParseUint(params["productId"], 10, 64); err != nil {
		log.Printf("Failed to parse path parameter, %v", err)
		w.WriteHeader(400)
		return
	}

	if p, err := pr.productService.FindById(id); err != nil {
		log.Printf("Failed to get product, %v", err)
		w.WriteHeader(500)
	} else if p == nil {
		w.WriteHeader(404)
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(p)
	}
}

func (pr *ProductRoute) Create(w http.ResponseWriter, r *http.Request) {

	var p model.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("Failed to parse request body, %v", err)
		w.WriteHeader(400)
		return
	}

	if err := pr.validateCreate(&p); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response.ErrorResponse{"BAD_REQUEST", fmt.Sprintf("%v", err)})
		return
	}

	persisted, err := pr.productService.Create(&p)
	if err != nil {
		log.Printf("Failed to create product, %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("http://%s%s/%d", r.Host, r.URL.Path, persisted.ID))
	w.WriteHeader(201)
}

func (pr *ProductRoute) validateCreate(p *model.Product) error {

	if p == nil {
		return fmt.Errorf("Request body cant be null")
	}
	if len(strings.TrimSpace(p.Name)) == 0 {
		return fmt.Errorf("Name property cant be empty")
	}
	if len(strings.TrimSpace(p.Type)) == 0 {
		return fmt.Errorf("Type property cant be empty")
	}
	if len(strings.TrimSpace(p.Description)) == 0 {
		return fmt.Errorf("Description property cant be empty")
	}
	if p.Price == 0 {
		return fmt.Errorf("Price property cant be zero")
	}
	return nil
}
