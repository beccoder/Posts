package Blogs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CreatePostRequest struct {
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}

type UpdatePostRequest struct {
	Title *string `json:"title"`
	Text  *string `json:"text"`
}

type PostResponse struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AuthorsId primitive.ObjectID `json:"authors_id" bson:"authors_id"`
	Title     string             `json:"title" bson:"title"`
	Text      string             `json:"text" bson:"text"`
	Likes     []LikeResponse     `json:"likes,omitempty" bson:"likes,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type LikeResponse struct {
	LikedById primitive.ObjectID `json:"liked_by_id" bson:"liked_by_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type CreateCommentRequest struct {
	ReplyPostId primitive.ObjectID `json:"reply_post_id,omitempty"`
	Comment     string             `json:"comment" binding:"required"`
}

type UpdateCommentRequest struct {
	ReplyPostId *primitive.ObjectID `json:"reply_post_id"`
	Comment     *string             `json:"comment"`
}

type CommentResponse struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PostId        primitive.ObjectID `json:"post_id" bson:"post_id"`
	CommentedById primitive.ObjectID `json:"commented_by_id" bson:"commented_by_id"`
	ReplyPostId   primitive.ObjectID `json:"reply_post_id,omitempty" bson:"reply_post_id,omitempty"`
	Comment       string             `json:"comment" bson:"comment"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     *time.Time         `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
