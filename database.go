package Blogs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserModel struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Role      string             `bson:"role"`
	FirstName string             `bson:"first_name" binding:"required"`
	LastName  string             `bson:"last_name,omitempty"`
	Username  string             `bson:"username" binding:"required"`
	Password  string             `bson:"password" binding:"required"`
	Email     string             `bson:"email" binding:"required"`
	Bio       string             `bson:"bio,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

type PostModel struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	AuthorsId primitive.ObjectID `bson:"authors_id"`
	Title     string             `bson:"title"`
	Text      string             `bson:"text"`
	Likes     []LikeModel        `bson:"likes,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

type LikeModel struct {
	LikedById primitive.ObjectID `bson:"liked_by_id"`
	CreatedAt time.Time          `bson:"created_at"`
}

type CommentModel struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	PostId        primitive.ObjectID `bson:"post_id"`
	CommentedById primitive.ObjectID `bson:"commented_by_id"`
	ReplyPostId   primitive.ObjectID `bson:"reply_post_id,omitempty"`
	Comment       string             `bson:"comment"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at,omitempty"`
}
