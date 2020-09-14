package config

import (
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/model"
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/service"
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/web/route"
)

func CreateProductRoute() *route.ProductRoute {
	pd := model.NewProductDao()
	ps := service.NewProductService(pd)
	pr := route.NewProductRoute(ps)
	return pr
}
