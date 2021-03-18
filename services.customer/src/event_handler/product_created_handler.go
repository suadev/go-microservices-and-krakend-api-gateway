package eventhandler

import (
	"encoding/json"

	"github.com/suadev/microservices/services.customer/src/internal"
	"github.com/suadev/microservices/services.customer/src/entity"
)

func CreateProduct(service *customer.Service, message []byte) {
	var product entity.Product
	json.Unmarshal(message, &product)
	service.CreateProduct(product)
}
