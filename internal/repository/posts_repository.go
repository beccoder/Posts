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

func (p *PostsRepo) CreatePosts(post Blogs.Post) (primitive.ObjectID, error) {
	post.CreatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}
	collPosts := p.db.Database("blogs").Collection("posts")

	result, err := collPosts.InsertOne(context.TODO(), post)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (p *PostsRepo) GetMyAllPosts(userId primitive.ObjectID) ([]Blogs.Post, error) {
	collPosts := p.db.Database("blogs").Collection("posts")
	result, err := collPosts.Find(context.TODO(), bson.M{"authors_id": userId})

	var posts []Blogs.Post
	if err = result.All(context.TODO(), &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostsRepo) GetAllPosts() ([]Blogs.Post, error) {
	collPosts := p.db.Database("blogs").Collection("posts")
	result, err := collPosts.Find(context.TODO(), bson.M{})

	var posts []Blogs.Post
	if err = result.All(context.TODO(), &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostsRepo) GetPostById(postId primitive.ObjectID) (Blogs.Post, error) {
	var post Blogs.Post
	collPosts := p.db.Database("blogs").Collection("posts")
	err := collPosts.FindOne(context.TODO(), bson.D{{"_id", postId}}).Decode(&post)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Blogs.Post{}, errors.New("no posts exist")
		}
		return Blogs.Post{}, err
	}
	return post, nil
}

func (p *PostsRepo) UpdatePost(postId primitive.ObjectID, input Blogs.PostUpdate) error {
	var update bson.D
	if input.Title != "" {
		update = append(update, bson.E{"$set", bson.D{{"title", input.Title}}})
	}

	if input.Text != "" {
		update = append(update, bson.E{"$set", bson.D{{"text", input.Text}}})
	}

	if len(update) == 0 {
		return errors.New("no fields to update")
	}

	update = append(update, bson.E{"$set", bson.D{{"updated_at", primitive.Timestamp{T: uint32(time.Now().Unix())}}}})
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

func (p *PostsRepo) CreateComment(input Blogs.Comment) (primitive.ObjectID, error) {
	input.CreatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}
	collPosts := p.db.Database("blogs").Collection("comments")

	result, err := collPosts.InsertOne(context.TODO(), input)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (p *PostsRepo) GetAllComments(postId primitive.ObjectID) ([]Blogs.Comment, error) {
	collPosts := p.db.Database("blogs").Collection("comments")
	result, err := collPosts.Find(context.TODO(), bson.M{"post_id": postId})

	var comments []Blogs.Comment
	if err = result.All(context.TODO(), &comments); err != nil {
		return nil, err
	}
	return comments, nil
}

func (p *PostsRepo) GetCommentById(commentId primitive.ObjectID) (Blogs.Comment, error) {
	var comment Blogs.Comment
	collPosts := p.db.Database("blogs").Collection("comments")
	err := collPosts.FindOne(context.TODO(), bson.D{{"_id", commentId}}).Decode(&comment)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Blogs.Comment{}, errors.New("no comments exist")
		}
		return Blogs.Comment{}, err
	}
	return comment, nil
}

func (p *PostsRepo) UpdateComment(commentId primitive.ObjectID, input Blogs.CommentUpdate) error {
	var update bson.D
	if input.Comment != "" {
		update = append(update, bson.E{"$set", bson.D{{"comment", input.Comment}}})
	}

	if !input.ReplyPostId.IsZero() {
		update = append(update, bson.E{"$set", bson.D{{"reply_post_id", input.ReplyPostId}}})
	}

	if len(update) == 0 {
		return errors.New("no fields to update")
	}

	update = append(update, bson.E{"$set", bson.D{{"updated_at", primitive.Timestamp{T: uint32(time.Now().Unix())}}}})
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
	update := bson.D{{"$push", bson.D{{"likes", Blogs.Like{
		LikedById: likedById,
		CreatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
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
