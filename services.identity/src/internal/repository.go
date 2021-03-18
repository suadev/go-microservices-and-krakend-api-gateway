package identity

import (
	"github.com/suadev/microservices/services.identity/src/entity"
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

func (r *Repository) IsUserExist(email string) bool {
	user := entity.User{}
	r.db.Find(&user, "email = ?", email)
	return user.Email != ""
}

func (r *Repository) Create(user *entity.User) (*entity.User, error) {
	err := r.db.Model(&entity.User{}).Create(&user).Error
	return user, err
}

func (r *Repository) GetUserByEmail(email string, password string) (entity.User, error) {
	var user entity.User
	err := r.db.Where(entity.User{Email: email, Password: password}).Take(&user).Error
	return user, err
}
