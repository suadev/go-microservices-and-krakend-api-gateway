package order

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) health(context *gin.Context) {
	context.JSON(http.StatusOK, "Order service is up!")
}

func (h *Handler) getOrders(context *gin.Context) {

	customers, err := h.service.GetOrders()
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, customers)
}

func (h *Handler) getOrder(context *gin.Context) {

	id, _ := uuid.Parse(context.Param("id"))
	order, err := h.service.GetOrder(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, order)
}

func (h *Handler) createOrder(context *gin.Context) {

	var userID = context.Request.Header["User_id"][0] // injected by KrakenD
	order, err := h.service.CreateOrder(userID)
	if err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, order)
}
