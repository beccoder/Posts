package repository

import (
	"Blogs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Administration interface {
	CreateUser(input Blogs.UserModel) (primitive.ObjectID, error)
	GetAllUsers() ([]Blogs.UserResponse, error)
	GetUserById(userId primitive.ObjectID) (Blogs.UserResponse, error)
	UpdateUser(userId primitive.ObjectID, input Blogs.UpdateUserRequest) error
	DeleteUser(userId primitive.ObjectID) error
}

type Authorization interface {
	//sign-in
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
	Administration
	Authorization
	Posts
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Administration: NewAdmRepo(db),
		Authorization:  NewAuthRepo(db),
		Posts:          NewPostsRepo(db),
	}
}
