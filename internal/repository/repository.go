package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	//sign-in
	//sign-up
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
