package service

import (
	"Blogs"
	"Blogs/internal/repository"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdmService struct {
	repo repository.Administration
}

func NewAdmService(repo repository.Administration) *AdmService {
	return &AdmService{repo: repo}
}

func (s *AdmService) CreateUser(input Blogs.UserModel) (primitive.ObjectID, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return primitive.ObjectID{}, err
	}
	for _, user := range users {
		if input.Email == user.Email && input.Role == user.Role {
			return primitive.ObjectID{}, errors.New("already registered")
		}
		if input.Username == user.Username {
			return primitive.ObjectID{}, errors.New("username exists")
		}
	}
	input.Password = Blogs.GeneratePasswordHash(input.Password)
	return s.repo.CreateUser(input)
}

func (s *AdmService) UpdateUser(userId primitive.ObjectID, input Blogs.UpdateUserRequest) error {
	if input.Password != nil {
		passwordHash := Blogs.GeneratePasswordHash(*input.Password)
		input.Password = &passwordHash
	}
	return s.repo.UpdateUser(userId, input)
}

func (s *AdmService) DeleteUser(userId primitive.ObjectID) error {
	return s.repo.DeleteUser(userId)
}

func (s *AdmService) GetAllUsers() ([]Blogs.UserResponse, error) {
	return s.repo.GetAllUsers()
}

func (s *AdmService) GetUserById(userId primitive.ObjectID) (Blogs.UserResponse, error) {
	return s.repo.GetUserById(userId)
}
