package service

import (
	"Blogs"
	"Blogs/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Administration interface {
	CreateUser(input Blogs.UserModel) (primitive.ObjectID, error)
	GetAllUsers() ([]Blogs.UserResponse, error)
	GetUserById(userId primitive.ObjectID) (Blogs.UserResponse, error)
	UpdateUser(userId primitive.ObjectID, input Blogs.UpdateUserRequest) error
	DeleteUser(userId primitive.ObjectID) error
}

type Authorization interface {
	GenerateToken(login, password, role string) (string, error)
	GetUserRole(login, password string) (string, error)
	ParseToken(accessToken string) (primitive.ObjectID, string, error)
}

type Posts interface {
	// CRUD posts
	CreatePosts(post Blogs.PostModel) (primitive.ObjectID, error)
	GetMyAllPosts(userId primitive.ObjectID) ([]Blogs.PostResponse, error)
	GetAllPosts() ([]Blogs.PostResponse, error)
	GetPostById(postId primitive.ObjectID) (Blogs.PostResponse, error)
	UpdatePost(postId primitive.ObjectID, input Blogs.UpdatePostRequest) error
	DeletePost(postId primitive.ObjectID) error
	// Comments
	CreateComment(input Blogs.CommentModel) (primitive.ObjectID, error)
	GetAllComments(postId primitive.ObjectID) ([]Blogs.CommentResponse, error)
	GetCommentById(commentId primitive.ObjectID) (Blogs.CommentResponse, error)
	UpdateComment(commentId primitive.ObjectID, input Blogs.UpdateCommentRequest) error
	DeleteComment(commentId primitive.ObjectID) error
	// Likes
	AddLike(postId primitive.ObjectID, likedById primitive.ObjectID) error
	UnlikePost(postId primitive.ObjectID, likedById primitive.ObjectID) error
}

type Service struct {
	Administration
	Authorization
	Posts
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Administration: NewAdmService(repos.Administration),
		Authorization:  NewAuthService(repos.Authorization),
		Posts:          NewPostsService(repos.Posts),
	}
}
