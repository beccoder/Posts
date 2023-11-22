package repository

import (
	"Blogs"
	"context"
	"errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepo struct {
	db *mongo.Client
}

func NewAuthRepo(db *mongo.Client) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) GetUser(username, password string) (Blogs.UserResponse, error) {
	var user Blogs.UserResponse
	coll := r.db.Database(viper.GetString("MONGO.DATABASE")).Collection("users")
	err := coll.FindOne(context.TODO(), bson.D{{"username", username}, {"password", password}}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Blogs.UserResponse{}, errors.New("invalid username or password")
		}
		return Blogs.UserResponse{}, err
	}
	return user, nil
}
