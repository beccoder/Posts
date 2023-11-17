package repository

import (
	"Blogs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	//sign-in
	CreateUser(input Blogs.UserModel) (primitive.ObjectID, error)
	GetAllUsers() ([]Blogs.UserResponse, error)
	GetUser(username, password string) (Blogs.UserResponse, error)
}

type Posts interface {
	// CRUD
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
	//Likes
	AddLike(postId primitive.ObjectID, likedById primitive.ObjectID) error
	UnlikePost(postId primitive.ObjectID, likedById primitive.ObjectID) error
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
