package order

import (
	"encoding/json"

	"github.com/google/uuid"
	gouuid "github.com/satori/go.uuid"
	"github.com/suadev/microservices/services.order/src/dto"
	"github.com/suadev/microservices/services.order/src/entity"
	"github.com/suadev/microservices/services.order/src/event"
	httpclient "github.com/suadev/microservices/services.order/src/http_client"
	"github.com/suadev/microservices/shared/config"
	"github.com/suadev/microservices/shared/kafka"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

//CreateOrder ...
func (s *Service) CreateOrder(customerID string) (entity.Order, error) {
	order := entity.Order{}
	order.ID = uuid.New()

	client := httpclient.NewCustomerClient()
	basketItems, err := client.GetBasketItems(customerID)
	if err != nil {
		return order, err
	}

	customerIDGouuid, _ := gouuid.FromString(customerID)
	order.Status = entity.OrderCreated
	order.CustomerID = uuid.UUID(customerIDGouuid)

	var orderItems []entity.OrderItem
	for _, basketItem := range basketItems {
		orderItem := entity.OrderItem{
			ID:          uuid.New(),
			OrderID:     order.ID,
			ProductID:   basketItem.ProductID,
			ProductName: basketItem.ProductName,
			UnitPrice:   basketItem.UnitPrice,
			Quantity:    basketItem.Quantity,
		}
		orderItems = append(orderItems, orderItem)

		order.TotalAmount += (basketItem.UnitPrice * float64(basketItem.Quantity))
	}

	createdOrder, _, err := s.repo.CreateOrder(order, orderItems)
	if err != nil {
		return createdOrder, err
	}

	publishOrderCreatedEvent(createdOrder, basketItems)

	return createdOrder, err
}

func publishOrderCreatedEvent(createdOrder entity.Order, basketItems []dto.BasketItemDto) {
	orderCreatedEvent := event.OrderCreated{
		ID:          createdOrder.ID,
		CustomerID:  createdOrder.CustomerID,
		TotalAmount: createdOrder.TotalAmount,
	}

	for _, basketItem := range basketItems {
		orderCreatedEvent.Items = append(orderCreatedEvent.Items,
			event.OrderBasketItem{ProductID: basketItem.ProductID,
				Quantity: basketItem.Quantity})
	}

	payload, _ := json.Marshal(orderCreatedEvent)
	kafka.Publish(createdOrder.ID, payload, "OrderCreated", config.AppConfig.KafkaOrderTopic)
}

func (s *Service) UpdateOrderStatus(id uuid.UUID, status int) (entity.Order, error) {
	return s.repo.UpdateOrderStatus(id, status)
}

func (s *Service) GetOrders() ([]entity.Order, error) {
	return s.repo.GetList()
}

func (s *Service) GetOrder(id uuid.UUID) (entity.Order, error) {
	return s.repo.GetById(id)
}
