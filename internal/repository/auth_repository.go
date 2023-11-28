package repository

import (
	"Blogs"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepo struct {
	client   *mongo.Client
	database string
}

func NewAuthRepo(client *mongo.Client, database string) *AuthRepo {
	return &AuthRepo{client: client, database: database}
}

func (r *AuthRepo) GetUser(username, password string) (Blogs.UserResponse, error) {
	var user Blogs.UserResponse
	coll := r.client.Database(r.database).Collection("users")
	err := coll.FindOne(context.TODO(), bson.D{{"username", username}, {"password", password}}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Blogs.UserResponse{}, errors.New("invalid username or password")
		}
		return Blogs.UserResponse{}, err
	}
	return user, nil
}
