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

type AdmRepo struct {
	client   *mongo.Client
	database string
}

func NewAdmRepo(client *mongo.Client, database string) *AdmRepo {
	return &AdmRepo{client: client, database: database}
}

func (r *AdmRepo) CreateUser(input Blogs.UserModel) (primitive.ObjectID, error) {
	input.CreatedAt = time.Now()
	collUsers := r.client.Database(r.database).Collection("users")

	result, err := collUsers.InsertOne(context.TODO(), input)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *AdmRepo) GetAllUsers() ([]Blogs.UserResponse, error) {
	var users []Blogs.UserResponse
	coll := r.client.Database(r.database).Collection("users")
	filter, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = filter.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *AdmRepo) GetUserById(userId primitive.ObjectID) (Blogs.UserResponse, error) {
	var user Blogs.UserResponse
	coll := r.client.Database(r.database).Collection("users")
	err := coll.FindOne(context.TODO(), bson.D{{"_id", userId}}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Blogs.UserResponse{}, errors.New("invalid id")
		}
		return Blogs.UserResponse{}, err
	}
	return user, nil
}

func (r *AdmRepo) UpdateUser(userId primitive.ObjectID, input Blogs.UpdateUserRequest) error {
	var update bson.D
	if input.FirstName != nil {
		update = append(update, bson.E{"$set", bson.D{{"first_name", *input.FirstName}}})
	}

	if input.LastName != nil {
		update = append(update, bson.E{"$set", bson.D{{"last_name", *input.LastName}}})
	}

	if input.Username != nil {
		update = append(update, bson.E{"$set", bson.D{{"username", *input.Username}}})
	}

	if input.Password != nil {
		update = append(update, bson.E{"$set", bson.D{{"password", *input.Password}}})
	}

	if input.Role != nil {
		update = append(update, bson.E{"$set", bson.D{{"role", *input.Role}}})
	}

	if input.Email != nil {
		update = append(update, bson.E{"$set", bson.D{{"email", *input.Email}}})
	}

	if input.Bio != nil {
		update = append(update, bson.E{"$set", bson.D{{"bio", *input.Bio}}})
	}

	if len(update) == 0 {
		return errors.New("no fields to update")
	}

	updatedTime := time.Now()
	update = append(update, bson.E{"$set", bson.D{{"updated_at", &updatedTime}}})
	collPosts := r.client.Database(r.database).Collection("users")

	res, err := collPosts.UpdateOne(context.TODO(), bson.D{{"_id", userId}}, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return errors.New("no matching user")
	}
	return nil
}

func (r *AdmRepo) DeleteUser(userId primitive.ObjectID) error {
	collPosts := r.client.Database(r.database).Collection("users")
	result, err := collPosts.DeleteOne(context.TODO(), bson.D{{"_id", userId}})

	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no matching user")
	}
	return nil
}
