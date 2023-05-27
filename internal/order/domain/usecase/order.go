package usecase

import "services/internal/order/repository/mgo"

type IOrderUseCase interface {
	CreateOrder()
}

type orderUseCase struct {
	orderRepo mgo.IOrderRepository
}

func NewOrderUseCase(orderRepo mgo.IOrderRepository) IOrderUseCase {
	return &orderUseCase{
		orderRepo: orderRepo,
	}
}

// b
func (o *orderUseCase) CreateOrder() {

}
