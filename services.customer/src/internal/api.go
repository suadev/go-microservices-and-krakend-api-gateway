package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/suadev/microservices/services.customer/src/entity"
)

func (h *Handler) health(context *gin.Context) {
	context.JSON(http.StatusOK, "Customer service is up!")
}

func (h *Handler) getCustomers(context *gin.Context) {

	customers, err := h.service.GetCustomers()
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	context.JSON(http.StatusOK, customers)
}

func (h *Handler) getCustomer(context *gin.Context) {

	id, _ := uuid.Parse(context.Param("id"))
	customer, err := h.service.GetCustomer(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	context.JSON(http.StatusOK, customer)
}

func (h *Handler) getBasketItems(context *gin.Context) {

	id, _ := uuid.Parse(context.Param("id"))
	basketItems, err := h.service.GetBasketItems(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	context.JSON(http.StatusOK, basketItems)
}

func (h *Handler) addItemToBasket(context *gin.Context) {

	var userID = context.Request.Header["User_id"][0] // injected by KrakenD
	// fmt.Println(userID)
	var input *entity.BasketItem
	context.BindJSON(&input)
	item, err := h.service.AddItemToBasket(input, userID)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, item)
}
