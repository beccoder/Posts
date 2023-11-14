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
	//Likes
	AddLike(postId primitive.ObjectID, likedById primitive.ObjectID) error
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
