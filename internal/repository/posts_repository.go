package repository

import "go.mongodb.org/mongo-driver/mongo"

type PostsRepo struct {
	db *mongo.Client
}

func NewPostsRepo(db *mongo.Client) *PostsRepo {
	return &PostsRepo{db: db}
}
