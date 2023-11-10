package repository

import (
	"Blogs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	//sign-in
	CreateUser(input Blogs.User) (primitive.ObjectID, error)
	GetAllUsers() ([]Blogs.User, error)
	GetUser(username, password string) (Blogs.User, error)
}

type Posts interface {
	// CRUD
	// Comments
}

type Repository struct {
	Authorization
	Posts
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
		Posts:         NewPostsRepo(db),
	}
}
