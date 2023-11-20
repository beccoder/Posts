package repository

import (
	"Blogs"
	"context"
	"errors"
	"github.com/spf13/viper"
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

func (r *AuthRepo) CreateUser(input Blogs.UserModel) (primitive.ObjectID, error) {
	input.CreatedAt = time.Now()
	collUsers := r.db.Database(viper.GetString("MONGO.DATABASE")).Collection("users")

	result, err := collUsers.InsertOne(context.TODO(), input)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *AuthRepo) GetAllUsers() ([]Blogs.UserResponse, error) {
	var users []Blogs.UserResponse
	coll := r.db.Database(viper.GetString("MONGO.DATABASE")).Collection("users")
	filter, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = filter.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *AuthRepo) GetUser(username, password string) (Blogs.UserResponse, error) {
	var user Blogs.UserResponse
	coll := r.db.Database(viper.GetString("MONGO.DATABASE")).Collection("users")
	err := coll.FindOne(context.TODO(), bson.D{{"username", username}, {"password", password}}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Blogs.UserResponse{}, errors.New("invalid username or password")
		}
		return Blogs.UserResponse{}, err
	}
	return user, nil
}
