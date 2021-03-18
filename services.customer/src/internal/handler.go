package customer

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
	routerGroup.GET("/customers", h.getCustomers)
	routerGroup.GET("/customers/:id", h.getCustomer)
	routerGroup.GET("/customers/:id/basketItems", h.getBasketItems)
	routerGroup.POST("/customer-basket", h.addItemToBasket)
}
