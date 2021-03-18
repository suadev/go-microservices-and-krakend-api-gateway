module github.com/suadev/microservices/services.product

go 1.15

replace github.com/suadev/microservices/shared => ../shared

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.2.0
	github.com/suadev/microservices/shared v0.0.0-00010101000000-000000000000
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.3
)
