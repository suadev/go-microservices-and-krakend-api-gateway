package product

import (
	"github.com/google/uuid"
	"github.com/suadev/microservices/services.product/src/entity"
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

func (r *Repository) Create(product *entity.Product) (*entity.Product, error) {
	err := r.db.Model(&product).Create(&product).Error
	return product, err
}

func (r *Repository) BulkUpdate(products *[]entity.Product) error {
	return r.db.Model(entity.Product{}).Save(products).Error
}

func (r *Repository) GetList() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *Repository) GetById(id uuid.UUID) (entity.Product, error) {
	var product entity.Product
	err := r.db.Where(entity.Product{ID: id}).Take(&product).Error
	return product, err
}
