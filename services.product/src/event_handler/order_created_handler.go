package eventhandler

import (
	"encoding/json"
	"log"

	"github.com/suadev/microservices/services.product/src/entity"
	"github.com/suadev/microservices/services.product/src/event"
	"github.com/suadev/microservices/services.product/src/internal"
	"github.com/suadev/microservices/shared/config"
	"github.com/suadev/microservices/shared/kafka"
)

func ReserveProducts(service *product.Service, message []byte) {

	isReserved := false
	var products []entity.Product

	var order event.OrderCreated
	json.Unmarshal(message, &order)

	for _, item := range order.Items {

		product, err := service.GetProduct(item.ProductID)
		if err != nil {
			log.Printf("Product not found: %v", item.ProductID.String())
			isReserved = false
			break
		}

		if item.Quantity > product.Quantity {
			isReserved = false
			log.Printf("Not available %v %v product.", item.Quantity, product.Name)
			break
		} else {
			isReserved = true
			product.Quantity -= item.Quantity
			products = append(products, product)
		}
	}

	eventType := "ProductReserveFailed"
	if isReserved {
		err := service.BulkUpdate(&products)
		if err == nil {
			eventType = "ProductReserved"
		}
	}
	kafka.Publish(order.ID, nil, eventType, config.AppConfig.KafkaProductTopic)
}
