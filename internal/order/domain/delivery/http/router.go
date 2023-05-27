package http

import "github.com/gin-gonic/gin"

type OrderRouter struct {
	orderHandler OrderHandler
}

func NewOrderRouter(handler OrderHandler) OrderRouter {
	return OrderRouter{
		orderHandler: handler,
	}
}
func (h OrderRouter) UseOrderRoute(r *gin.RouterGroup) {
	group := r.Group("/order")
	group.GET("", h.orderHandler.ListOrder)
	group.POST("", h.orderHandler.CreateOrder)
}
