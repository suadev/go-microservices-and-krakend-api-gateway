package customer

import (
	"github.com/google/uuid"
	gouuid "github.com/satori/go.uuid"
	"github.com/suadev/microservices/services.customer/src/entity"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateCustomer(customer entity.Customer) (uuid.UUID, error) {
	return s.repo.CreateCustomer(customer)
}

func (s *Service) CreateProduct(product entity.Product) (uuid.UUID, error) {
	return s.repo.CreateProduct(product)
}

func (s *Service) GetCustomers() ([]entity.Customer, error) {
	return s.repo.GetList()
}

func (s *Service) GetCustomer(id uuid.UUID) (entity.Customer, error) {
	return s.repo.GetById(id)
}

func (s *Service) GetBasket(customerID uuid.UUID) (entity.Basket, error) {
	return s.repo.GetBasket(customerID)
}

func (s *Service) GetBasketItems(customerID uuid.UUID) ([]entity.BasketItem, error) {
	return s.repo.GetBasketItems(customerID)
}

func (s *Service) AddItemToBasket(basketItem *entity.BasketItem, userID string) (*entity.BasketItem, error) {

	customerID, _ := gouuid.FromString(userID)
	basket, err := s.repo.GetBasket(uuid.UUID(customerID))
	if err != nil {
		return basketItem, err
	}

	product, err := s.repo.GetProductById(basketItem.ProductID)
	if err != nil {
		return basketItem, err
	}

	existingBasketItem, err := s.repo.GetBasketItem(basketItem.ProductID)
	if err == nil {
		return s.repo.UpdateItemQuantity(&existingBasketItem, basketItem.Quantity)
	}

	basketItem.ID = uuid.New()
	basketItem.BasketID = basket.ID
	basketItem.ProductName = product.Name
	basketItem.UnitPrice = product.Price
	return s.repo.CreateBasketItem(basketItem)
}

func (s *Service) ClearBasketItems(basketID uuid.UUID) error {
	return s.repo.ClearBasketItems(basketID)
}
