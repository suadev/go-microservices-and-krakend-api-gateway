package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/suadev/microservices/services.product/src/entity"
)

func (h *Handler) health(context *gin.Context) {
	context.JSON(http.StatusOK, "Product service is up!")
}

func (h *Handler) getProducts(context *gin.Context) {

	products, err := h.service.GetProducts()
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	context.JSON(http.StatusOK, products)
}

func (h *Handler) getProduct(context *gin.Context) {

	id, _ := uuid.Parse(context.Param("id"))
	product, err := h.service.GetProduct(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	context.JSON(http.StatusOK, product)
}

func (h *Handler) createProduct(context *gin.Context) {

	var input *entity.Product
	context.BindJSON(&input)
	product, err := h.service.CreateProduct(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, product)
}
