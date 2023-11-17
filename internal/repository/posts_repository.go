package repository

import (
	"Blogs"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type PostsRepo struct {
	db *mongo.Client
}

func NewPostsRepo(db *mongo.Client) *PostsRepo {
	return &PostsRepo{db: db}
}

func (p *PostsRepo) CreatePosts(post Blogs.PostModel) (primitive.ObjectID, error) {
	post.CreatedAt = time.Now()
	collPosts := p.db.Database("blogs").Collection("posts")

	result, err := collPosts.InsertOne(context.TODO(), post)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (p *PostsRepo) GetMyAllPosts(userId primitive.ObjectID) ([]Blogs.PostResponse, error) {
	collPosts := p.db.Database("blogs").Collection("posts")
	result, err := collPosts.Find(context.TODO(), bson.M{"authors_id": userId})

	var posts []Blogs.PostResponse
	if err = result.All(context.TODO(), &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostsRepo) GetAllPosts() ([]Blogs.PostResponse, error) {
	collPosts := p.db.Database("blogs").Collection("posts")
	result, err := collPosts.Find(context.TODO(), bson.M{})

	var posts []Blogs.PostResponse
	if err = result.All(context.TODO(), &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostsRepo) GetPostById(postId primitive.ObjectID) (Blogs.PostResponse, error) {
	var post Blogs.PostResponse
	collPosts := p.db.Database("blogs").Collection("posts")
	err := collPosts.FindOne(context.TODO(), bson.D{{"_id", postId}}).Decode(&post)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Blogs.PostResponse{}, errors.New("no posts exist")
		}
		return Blogs.PostResponse{}, err
	}
	return post, nil
}

func (p *PostsRepo) UpdatePost(postId primitive.ObjectID, input Blogs.UpdatePostRequest) error {
	var update bson.D
	if input.Title != nil {
		update = append(update, bson.E{"$set", bson.D{{"title", *input.Title}}})
	}

	if input.Text != nil {
		update = append(update, bson.E{"$set", bson.D{{"text", *input.Text}}})
	}

	if len(update) == 0 {
		return errors.New("no fields to update")
	}

	update = append(update, bson.E{"$set", bson.D{{"updated_at", time.Now()}}})
	collPosts := p.db.Database("blogs").Collection("posts")

	_, err := collPosts.UpdateOne(context.TODO(), bson.D{{"_id", postId}}, update)
	return err
}

func (p *PostsRepo) DeletePost(postId primitive.ObjectID) error {
	collPosts := p.db.Database("blogs").Collection("posts")
	result, err := collPosts.DeleteOne(context.TODO(), bson.D{{"_id", postId}})

	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no posts deleted")
	}
	return nil
}

func (p *PostsRepo) CreateComment(input Blogs.CommentModel) (primitive.ObjectID, error) {
	input.CreatedAt = time.Now()
	collPosts := p.db.Database("blogs").Collection("comments")

	result, err := collPosts.InsertOne(context.TODO(), input)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (p *PostsRepo) GetAllComments(postId primitive.ObjectID) ([]Blogs.CommentResponse, error) {
	collPosts := p.db.Database("blogs").Collection("comments")
	result, err := collPosts.Find(context.TODO(), bson.M{"post_id": postId})

	var comments []Blogs.CommentResponse
	if err = result.All(context.TODO(), &comments); err != nil {
		return nil, err
	}
	return comments, nil
}

func (p *PostsRepo) GetCommentById(commentId primitive.ObjectID) (Blogs.CommentResponse, error) {
	var comment Blogs.CommentResponse
	collPosts := p.db.Database("blogs").Collection("comments")
	err := collPosts.FindOne(context.TODO(), bson.D{{"_id", commentId}}).Decode(&comment)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Blogs.CommentResponse{}, errors.New("no comments exist")
		}
		return Blogs.CommentResponse{}, err
	}
	return comment, nil
}

func (p *PostsRepo) UpdateComment(commentId primitive.ObjectID, input Blogs.UpdateCommentRequest) error {
	var update bson.D
	if input.Comment != nil {
		update = append(update, bson.E{"$set", bson.D{{"comment", *input.Comment}}})
	}

	if input.ReplyPostId != nil {
		update = append(update, bson.E{"$set", bson.D{{"reply_post_id", *input.ReplyPostId}}})
	}

	if len(update) == 0 {
		return errors.New("no fields to update")
	}

	update = append(update, bson.E{"$set", bson.D{{"updated_at", time.Now()}}})
	collComments := p.db.Database("blogs").Collection("comments")

	_, err := collComments.UpdateOne(context.TODO(), bson.D{{"_id", commentId}}, update)
	return err
}

func (p *PostsRepo) DeleteComment(commentId primitive.ObjectID) error {
	collComments := p.db.Database("blogs").Collection("comments")
	result, err := collComments.DeleteOne(context.TODO(), bson.D{{"_id", commentId}})

	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no comments deleted")
	}
	return nil
}

func (p *PostsRepo) AddLike(postId primitive.ObjectID, likedById primitive.ObjectID) error {
	post, err := p.GetPostById(postId)
	if err != nil {
		return err
	}
	if post.Likes != nil {
		for _, like := range post.Likes {
			if like.LikedById == likedById {
				return errors.New("already liked")
			}
		}
	}

	filter := bson.D{{"_id", postId}}
	update := bson.D{{"$push", bson.D{{"likes", Blogs.LikeModel{
		LikedById: likedById,
		CreatedAt: time.Now(),
	}}}}}
	collPosts := p.db.Database("blogs").Collection("posts")

	result, err := collPosts.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("like is not added")
	}
	return nil
}

func (p *PostsRepo) UnlikePost(postId primitive.ObjectID, likedById primitive.ObjectID) error {
	filter := bson.D{{"_id", postId}}
	update := bson.D{{"$pull", bson.D{{"likes", bson.D{{"liked_by_id", likedById}}}}}}

	collPosts := p.db.Database("blogs").Collection("posts")

	_, err := collPosts.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
