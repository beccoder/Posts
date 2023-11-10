package service

import (
	"Blogs"
	"Blogs/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Authorization interface {
	//sign-in
	GenerateToken(login, password string) (string, error)
	CreateUser(input Blogs.User) (primitive.ObjectID, error)
	GetUserRole(login, password string) (string, error)
	ParseToken(accessToken string) (primitive.ObjectID, string, error)
}

type Posts interface {
	// CRUD
	// Comments
}

type Service struct {
	Authorization
	Posts
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Posts:         NewPostsService(repos.Posts),
	}
}
