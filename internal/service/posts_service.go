package service

import (
	"Blogs"
	"Blogs/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostsService struct {
	repo repository.Posts
}

func NewPostsService(repo repository.Posts) *PostsService {
	return &PostsService{repo: repo}
}

func (p *PostsService) CreatePosts(post Blogs.PostModel) (primitive.ObjectID, error) {
	return p.repo.CreatePosts(post)
}

func (p *PostsService) GetMyAllPosts(userId primitive.ObjectID) ([]Blogs.PostResponse, error) {
	return p.repo.GetMyAllPosts(userId)
}

func (p *PostsService) GetAllPosts() ([]Blogs.PostResponse, error) {
	return p.repo.GetAllPosts()
}

func (p *PostsService) GetPostById(postId primitive.ObjectID) (Blogs.PostResponse, error) {
	return p.repo.GetPostById(postId)
}

func (p *PostsService) UpdatePost(postId primitive.ObjectID, input Blogs.UpdatePostRequest) error {
	return p.repo.UpdatePost(postId, input)
}

func (p *PostsService) DeletePost(postId primitive.ObjectID) error {
	return p.repo.DeletePost(postId)
}

func (p *PostsService) CreateComment(input Blogs.CommentModel) (primitive.ObjectID, error) {
	return p.repo.CreateComment(input)
}

func (p *PostsService) GetAllComments(postId primitive.ObjectID) ([]Blogs.CommentResponse, error) {
	return p.repo.GetAllComments(postId)
}

func (p *PostsService) GetCommentById(commentId primitive.ObjectID) (Blogs.CommentResponse, error) {
	return p.repo.GetCommentById(commentId)
}

func (p *PostsService) UpdateComment(commentId primitive.ObjectID, input Blogs.UpdateCommentRequest) error {
	return p.repo.UpdateComment(commentId, input)
}

func (p *PostsService) DeleteComment(commentId primitive.ObjectID) error {
	return p.repo.DeleteComment(commentId)
}

func (p *PostsService) AddLike(postId primitive.ObjectID, likedById primitive.ObjectID) error {
	return p.repo.AddLike(postId, likedById)
}

func (p *PostsService) UnlikePost(postId primitive.ObjectID, likedById primitive.ObjectID) error {
	return p.repo.UnlikePost(postId, likedById)
}
