package service

import (
	"Blogs"
	"Blogs/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Authorization interface {
	//sign-in
	GenerateToken(login, password, role string) (string, error)
	CreateUser(input Blogs.User) (primitive.ObjectID, error)
	GetUserRole(login, password string) (string, error)
	ParseToken(accessToken string) (primitive.ObjectID, string, error)
}

type Posts interface {
	// CRUD
	CreatePosts(post Blogs.Post) (primitive.ObjectID, error)
	GetMyAllPosts(userId primitive.ObjectID) ([]Blogs.Post, error)
	GetAllPosts() ([]Blogs.Post, error)
	GetPostById(postId primitive.ObjectID) (Blogs.Post, error)
	UpdatePost(postId primitive.ObjectID, input Blogs.PostUpdate) error
	DeletePost(postId primitive.ObjectID) error
	// Comments
	CreateComment(input Blogs.Comment) (primitive.ObjectID, error)
	GetAllComments(postId primitive.ObjectID) ([]Blogs.Comment, error)
	GetCommentById(commentId primitive.ObjectID) (Blogs.Comment, error)
	UpdateComment(commentId primitive.ObjectID, input Blogs.CommentUpdate) error
	DeleteComment(commentId primitive.ObjectID) error
	// Likes
	AddLike(postId primitive.ObjectID, likedById primitive.ObjectID) error
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
