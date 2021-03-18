package customer

import (
	"github.com/google/uuid"
	"github.com/suadev/microservices/services.customer/src/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	db.Logger.LogMode(logger.Info)
	return &Repository{db: db}
}

func (r *Repository) CreateCustomer(customer entity.Customer) (uuid.UUID, error) {

	err := r.db.Model(&entity.Customer{}).Create(&customer).Error
	if err == nil { // create empty basket for the new customer
		basket := entity.Basket{ID: uuid.New(), CustomerID: customer.ID}
		err = r.db.Model(&entity.Basket{}).Create(&basket).Error
	}
	return customer.ID, err
}

func (r *Repository) CreateProduct(product entity.Product) (uuid.UUID, error) {
	err := r.db.Model(&product).Create(&product).Error
	return product.ID, err
}

func (r *Repository) GetList() ([]entity.Customer, error) {
	var customers []entity.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r *Repository) GetById(id uuid.UUID) (entity.Customer, error) {
	var customer entity.Customer
	err := r.db.Where(entity.Customer{ID: id}).Take(&customer).Error
	return customer, err
}

func (r *Repository) GetProductById(id uuid.UUID) (entity.Product, error) {
	var product entity.Product
	err := r.db.Where(entity.Product{ID: id}).Take(&product).Error
	return product, err
}

func (r *Repository) GetBasket(customerID uuid.UUID) (entity.Basket, error) {
	var basket entity.Basket
	err := r.db.Where(&entity.Basket{CustomerID: customerID}).Take(&basket).Error
	return basket, err
}

func (r *Repository) GetBasketItem(productID uuid.UUID) (entity.BasketItem, error) {
	var basketItem entity.BasketItem
	err := r.db.Where(&entity.BasketItem{ProductID: productID}).Take(&basketItem).Error
	return basketItem, err
}

func (r *Repository) GetBasketItems(customerID uuid.UUID) ([]entity.BasketItem, error) {
	basket, _ := r.GetBasket(customerID)
	var basketItems []entity.BasketItem
	err := r.db.Where(&entity.BasketItem{BasketID: basket.ID}).Find(&basketItems).Error
	return basketItems, err
}

func (r *Repository) CreateBasketItem(item *entity.BasketItem) (*entity.BasketItem, error) {
	err := r.db.Model(&entity.BasketItem{}).Create(&item).Error
	return item, err
}

func (r *Repository) UpdateItemQuantity(item *entity.BasketItem, quantity int) (*entity.BasketItem, error) {
	err := r.db.Model(&item).
		Where("id = ?", item.ID).Update("quantity", item.Quantity+quantity).Error
	return item, err
}

func (r *Repository) ClearBasketItems(basketID uuid.UUID) error {
	return r.db.Where(&entity.BasketItem{BasketID: basketID}).Delete(&entity.BasketItem{}).Error
}
