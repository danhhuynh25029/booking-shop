package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"service/services/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user models.User) error
}

type userRepository struct {
	mgo *mongo.Collection
}

func NewUserRepository(mgo *mongo.Collection) UserRepository {
	return userRepository{
		mgo: mgo,
	}
}
func (u userRepository) InsertUser(ctx context.Context, user models.User) error {
	_, err := u.mgo.InsertOne(ctx, user)
	return err
}
