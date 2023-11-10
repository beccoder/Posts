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

type AuthRepo struct {
	db *mongo.Client
}

func NewAuthRepo(db *mongo.Client) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) CreateUser(input Blogs.User) (primitive.ObjectID, error) {
	input.CreatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}
	coll := r.db.Database("blogs").Collection("users")

	result, err := coll.InsertOne(context.TODO(), input)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *AuthRepo) GetAllUsers() ([]Blogs.User, error) {
	var users []Blogs.User
	coll := r.db.Database("blogs").Collection("users")
	filter, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = filter.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *AuthRepo) GetUser(username, password string) (Blogs.User, error) {
	var user Blogs.User
	coll := r.db.Database("blogs").Collection("users")
	err := coll.FindOne(context.TODO(), bson.D{{"username", username}, {"password", password}}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// can handle if needed
			return Blogs.User{}, err
		}
		return Blogs.User{}, err
	}
	return user, nil
}
