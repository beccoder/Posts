package Blogs

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Role      string              `json:"role" bson:"role"`
	FirstName string              `json:"first_name" bson:"first_name"`
	LastName  string              `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Username  string              `json:"username" bson:"username"`
	Password  string              `json:"password" bson:"password"`
	Email     string              `json:"email" bson:"email"`
	Bio       string              `json:"bio,omitempty" bson:"bio,omitempty"`
	CreatedAt primitive.Timestamp `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.Timestamp `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt primitive.Timestamp `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}
