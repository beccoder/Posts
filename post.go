package Blogs

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostUpdate struct {
	Title string `json:"title" bson:"title"`
	Text  string `json:"text" bson:"text"`
}

type Post struct {
	Id        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	AuthorsId primitive.ObjectID  `json:"authors_id" bson:"authors_id"`
	Title     string              `json:"title" bson:"title" binding:"required"`
	Text      string              `json:"text" bson:"text" binding:"required"`
	Likes     []Like              `json:"likes,omitempty" bson:"likes,omitempty"`
	CreatedAt primitive.Timestamp `json:"-" bson:"created_at,omitempty"`
	UpdatedAt primitive.Timestamp `json:"-" bson:"updated_at,omitempty"`
}

type Like struct {
	LikedById primitive.ObjectID  `json:"liked_by_id" bson:"liked_by_id" binding:"required"`
	CreatedAt primitive.Timestamp `json:"-" bson:"created_at,omitempty"`
}

type CommentUpdate struct {
	ReplyPostId primitive.ObjectID `json:"reply_post_id,omitempty" bson:"reply_post_id,omitempty"`
	Comment     string             `json:"comment" bson:"comment" binding:"required"`
}

type Comment struct {
	Id            primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	PostId        primitive.ObjectID  `json:"post_id" bson:"post_id"`
	CommentedById primitive.ObjectID  `json:"commented_by_id" bson:"commented_by_id"`
	ReplyPostId   primitive.ObjectID  `json:"reply_post_id,omitempty" bson:"reply_post_id,omitempty"`
	Comment       string              `json:"comment" bson:"comment" binding:"required"`
	CreatedAt     primitive.Timestamp `json:"-" bson:"created_at,omitempty"`
	UpdatedAt     primitive.Timestamp `json:"-" bson:"updated_at,omitempty"`
}
