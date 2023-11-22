package service

import (
	"Blogs"
	"Blogs/internal/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	salt       = "dflshfksjdhsasdajc"
	signingKey = "dsjfhsiuesfsygfs437ds"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId primitive.ObjectID `json:"user_id"`
	Role   string             `json:"role"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) CreateUser(input Blogs.UserModel) (primitive.ObjectID, error) {
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
	input.Password = generatePasswordHash(input.Password)
	return s.repo.CreateUser(input)
}

func (s *AuthService) GenerateToken(login, password, role string) (string, error) {
	user, err := s.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	if user.Role != role {
		return "", errors.New("Invalid role")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Role,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) GetUserRole(login, password string) (string, error) {
	user, err := s.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	return user.Role, nil
}

func (s *AuthService) ParseToken(accessToken string) (primitive.ObjectID, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return primitive.ObjectID{}, "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return primitive.ObjectID{}, "", errors.New("Token claims are not of type *tokenClaims")
	}

	return claims.UserId, claims.Role, nil
}

func (s *AuthService) UpdateUser(userId primitive.ObjectID, input Blogs.UpdateUserRequest) error {
	if input.Password != nil {
		passwordHash := generatePasswordHash(*input.Password)
		input.Password = &passwordHash
	}
	return s.repo.UpdateUser(userId, input)
}

func (s *AuthService) DeleteUser(userId primitive.ObjectID) error {
	return s.repo.DeleteUser(userId)
}

func (s *AuthService) GetAllUsers() ([]Blogs.UserResponse, error) {
	return s.repo.GetAllUsers()
}

func (s *AuthService) GetUserById(userId primitive.ObjectID) (Blogs.UserResponse, error) {
	return s.repo.GetUserById(userId)
}
