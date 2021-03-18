package order

import (
	"github.com/google/uuid"
	"github.com/suadev/microservices/services.order/src/entity"
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

func (r *Repository) CreateOrder(order entity.Order, items []entity.OrderItem) (entity.Order, []entity.OrderItem, error) {

	err := r.db.Model(&entity.Order{}).Create(&order).Error
	if err != nil {
		return order, nil, err
	}
	r.db.Model(&[]entity.OrderItem{}).Create(&items)
	return order, items, err
}

func (r *Repository) UpdateOrderStatus(id uuid.UUID, status int) (entity.Order, error) {

	order, err := r.GetById(id)
	if err != nil {
		return order, err
	}
	err = r.db.Model(&order).Update("status", status).Error
	return order, err
}

func (r *Repository) GetList() ([]entity.Order, error) {
	var customers []entity.Order
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r *Repository) GetById(id uuid.UUID) (entity.Order, error) {
	order := &entity.Order{ID: id}
	err := r.db.Find(&order).Error
	return *order, err
}
