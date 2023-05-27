package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"services/internal/order/domain/usecase"
)

type OrderHandler struct {
	orderUseCase usecase.IOrderUseCase
}

func NewOrderHandler(orderUseCase usecase.IOrderUseCase) OrderHandler {
	return OrderHandler{
		orderUseCase: orderUseCase,
	}
}

func (h OrderHandler) CreateOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "ok"})
}

func (h OrderHandler) ListOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "ok"})
}
