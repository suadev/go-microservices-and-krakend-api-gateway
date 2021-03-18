package product

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	h.initRoutes(router)
	return router
}

func (h *Handler) initRoutes(router *gin.Engine) {
	routerGroup := router.Group("/api")
	routerGroup.GET("/health", h.health)
	routerGroup.GET("/products", h.getProducts)
	routerGroup.GET("/products/:id", h.getProduct)
	routerGroup.POST("/products", h.createProduct)
}
