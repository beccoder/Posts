package Blogs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserResponse struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Role      string             `json:"role" bson:"role"`
	FirstName string             `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Username  string             `json:"username" bson:"username" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required"`
	Bio       string             `json:"bio,omitempty" bson:"bio,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CreateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Bio       string `json:"bio,omitempty"`
}

type UpdateUserRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Username  *string `json:"username"`
	Password  *string `json:"password"`
	Email     *string `json:"email"`
	Role      *string `json:"role"`
	Bio       *string `json:"bio"`
}
