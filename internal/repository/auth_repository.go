package repository

import "go.mongodb.org/mongo-driver/mongo"

type AuthRepo struct {
	db *mongo.Client
}

func NewAuthRepo(db *mongo.Client) *AuthRepo {
	return &AuthRepo{db: db}
}
