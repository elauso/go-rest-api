package messaging

import (
	"encoding/json"
	"log"

	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/driver"
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/model"
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/service"
)

type ProductCreateConsumer struct {
	subscriber     *driver.Subscriber
	productService *service.ProductService
}

func NewProductCreateConsumer(ps *service.ProductService) *ProductCreateConsumer {
	pcc := &ProductCreateConsumer{}
	s := driver.NewSubscriber("productCreateChannel", pcc.Consume)
	pcc.subscriber = s
	pcc.productService = ps
	return pcc
}

func (pcc *ProductCreateConsumer) Consume(ch string, msg string) {
	log.Printf("Receiving product create event on channel[%s]: %s ", ch, msg)
	p := &model.Product{}
	if err := json.Unmarshal([]byte(msg), &p); err != nil {
		log.Printf("Failed to parse create product event, %v", err)
		return
	}
	if persisted, err := pcc.productService.Create(p); err != nil {
		log.Printf("Failed to persist create product event, %v", err)
	} else {
		log.Printf("Received product create event registered with success: %v", persisted)
	}
}
