package messaging

import (
	"encoding/json"
	"log"

	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/driver/pubsub"
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/model"
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/service"
)

type ProductCreateConsumer struct {
	subscriber     *pubsub.Subscriber
	productService *service.ProductService
}

func NewProductCreateConsumer(ps *service.ProductService) *ProductCreateConsumer {
	pcc := &ProductCreateConsumer{}
	if subscriber, err := pubsub.NewSubscriber("productCreateChannel", pcc.Listen); err != nil {
		panic(err)
	} else {
		pcc.subscriber = subscriber
		pcc.productService = ps
		return pcc
	}
}

func (pcc *ProductCreateConsumer) Listen(channel string, message string) {
	log.Printf("Receiving product create event on channel[%s]: %s ", channel, message)
	p := &model.Product{}
	if err := json.Unmarshal([]byte(message), &p); err != nil {
		log.Printf("Failed to parse create product event, %v", err)
	}
	if persisted, err := pcc.productService.Create(p); err != nil {
		log.Printf("Failed to persist create product event, %v", err)
	} else {
		log.Printf("Received product create event registered with success: %v", persisted)
	}
}
