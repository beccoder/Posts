package service

import (
	"Blogs/internal/repository"
)

type Authorization interface {
	//sign-in
	//sign-up
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
