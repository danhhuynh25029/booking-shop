package mgo

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type IOrderRepository interface {
	CreateOrder()
}

type orderRepository struct {
	collection *mongo.Collection
	redis      *redis.Client
}

func NewOrderRepository() IOrderRepository {
	return &orderRepository{}
}

func (o *orderRepository) CreateOrder() {

}
