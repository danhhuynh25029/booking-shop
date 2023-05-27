package init

import (
	"github.com/gin-gonic/gin"
	"services/internal/order/domain/delivery/http"
	"services/internal/order/domain/usecase"
	"services/internal/order/repository/mgo"
)

func init() {

}

func StartOrderService() error {
	r := gin.Default()
	group := r.Group("/v1")
	orderRepo := mgo.NewOrderRepository()
	orderUseCase := usecase.NewOrderUseCase(orderRepo)
	orderHandler := http.NewOrderHandler(orderUseCase)
	orderRouter := http.NewOrderRouter(orderHandler)
	orderRouter.UseOrderRoute(group)
	if err := r.Run(":8080"); err != nil {
		return err
	}
	return nil
}
