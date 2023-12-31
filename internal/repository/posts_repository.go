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
	client   *mongo.Client
	database string
}

func NewPostsRepo(client *mongo.Client, database string) *PostsRepo {
	return &PostsRepo{client: client, database: database}
}

func (p *PostsRepo) CreatePosts(post Blogs.PostModel) (primitive.ObjectID, error) {
	post.CreatedAt = time.Now()
	collPosts := p.client.Database(p.database).Collection("posts")

	result, err := collPosts.InsertOne(context.TODO(), post)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (p *PostsRepo) GetMyAllPosts(userId primitive.ObjectID) ([]Blogs.PostResponse, error) {
	collPosts := p.client.Database(p.database).Collection("posts")
	result, err := collPosts.Find(context.TODO(), bson.M{"authors_id": userId})

	var posts []Blogs.PostResponse
	if err = result.All(context.TODO(), &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostsRepo) GetAllPosts() ([]Blogs.PostResponse, error) {
	collPosts := p.client.Database(p.database).Collection("posts")
	result, err := collPosts.Find(context.TODO(), bson.M{})

	var posts []Blogs.PostResponse
	if err = result.All(context.TODO(), &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostsRepo) GetPostById(postId primitive.ObjectID) (Blogs.PostResponse, error) {
	var post Blogs.PostResponse
	collPosts := p.client.Database(p.database).Collection("posts")
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

	updatedTime := time.Now()
	update = append(update, bson.E{"$set", bson.D{{"updated_at", &updatedTime}}})
	collPosts := p.client.Database(p.database).Collection("posts")

	_, err := collPosts.UpdateOne(context.TODO(), bson.D{{"_id", postId}}, update)
	return err
}

func (p *PostsRepo) DeletePost(postId primitive.ObjectID) error {
	collPosts := p.client.Database(p.database).Collection("posts")
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
	collPosts := p.client.Database(p.database).Collection("comments")

	result, err := collPosts.InsertOne(context.TODO(), input)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (p *PostsRepo) GetAllComments(postId primitive.ObjectID) ([]Blogs.CommentResponse, error) {
	collPosts := p.client.Database(p.database).Collection("comments")
	result, err := collPosts.Find(context.TODO(), bson.M{"post_id": postId})

	var comments []Blogs.CommentResponse
	if err = result.All(context.TODO(), &comments); err != nil {
		return nil, err
	}
	return comments, nil
}

func (p *PostsRepo) GetCommentById(commentId primitive.ObjectID) (Blogs.CommentResponse, error) {
	var comment Blogs.CommentResponse
	collPosts := p.client.Database(p.database).Collection("comments")
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

	if input.ReplyCommentId != nil {
		update = append(update, bson.E{"$set", bson.D{{"reply_post_id", *input.ReplyCommentId}}})
	}

	if len(update) == 0 {
		return errors.New("no fields to update")
	}

	updatedTime := time.Now()
	update = append(update, bson.E{"$set", bson.D{{"updated_at", &updatedTime}}})
	collComments := p.client.Database(p.database).Collection("comments")

	res, err := collComments.UpdateOne(context.TODO(), bson.D{{"_id", commentId}}, update)
	if res.MatchedCount == 0 {
		return errors.New("invalid input, no matching id")
	}

	if res.ModifiedCount == 0 {
		return errors.New("nothing updated")
	}

	return err
}

func (p *PostsRepo) DeleteComment(commentId primitive.ObjectID) error {
	collComments := p.client.Database(p.database).Collection("comments")
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
	collPosts := p.client.Database(p.database).Collection("posts")

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
	post, err := p.GetPostById(postId)
	if err != nil {
		return err
	}

	liked := false
	if post.Likes != nil {
		for _, like := range post.Likes {
			if like.LikedById == likedById {
				liked = true
			}
		}
	}
	if liked == false {
		return errors.New("already unliked")
	}

	filter := bson.D{{"_id", postId}}
	update := bson.D{{"$pull", bson.D{{"likes", bson.D{{"liked_by_id", likedById}}}}}}

	collPosts := p.client.Database(p.database).Collection("posts")

	res, err := collPosts.UpdateOne(context.TODO(), filter, update)
	if res.MatchedCount == 0 {
		return errors.New("invalid post id")
	}
	if err != nil {
		return err
	}
	return nil
}
