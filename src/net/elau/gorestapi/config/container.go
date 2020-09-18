package config

import (
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/event/messaging"
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/model"
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/service"
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/web/route"
)

func CreateProductRoute() (*route.ProductRoute, *messaging.ProductCreateConsumer) {
	pd := model.NewProductDao()
	ps := service.NewProductService(pd)
	pcc := messaging.NewProductCreateConsumer(ps)
	pr := route.NewProductRoute(ps)
	return pr, pcc
}
