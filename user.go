package Blogs

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Role      string              `json:"role" bson:"role"`
	FirstName string              `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string              `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Username  string              `json:"username" bson:"username" binding:"required"`
	Password  string              `json:"password" bson:"password" binding:"required"`
	Email     string              `json:"email" bson:"email" binding:"required"`
	Bio       string              `json:"bio,omitempty" bson:"bio,omitempty"`
	CreatedAt primitive.Timestamp `json:"-" bson:"created_at,omitempty"`
	UpdatedAt primitive.Timestamp `json:"-" bson:"updated_at,omitempty"`
}
