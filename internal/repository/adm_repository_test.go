package repository

import (
	"Blogs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestAdmRepo_CreateUser(t *testing.T) {
	type fields struct {
		client   *mongo.Client
		database string
	}
	type args struct {
		input Blogs.UserModel
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    primitive.ObjectID
		wantErr bool
	}{
		{
			"CreateUserTestAdminSuccess1",
			fields{client: testClient, database: database},
			args{input: Blogs.UserModel{
				Id:        userIdList[0],
				Role:      "admin",
				FirstName: "Bekhzod",
				LastName:  "Khudoyarov",
				Username:  "beccoder",
				Password:  Blogs.GeneratePasswordHash("qwerty"),
				Email:     "bekhzodkhudoyarov@gmail.com",
				Bio:       "I am admin",
			}},
			userIdList[0],
			false,
		},
		{
			"CreateUserTestAdminFailDuplicateUsername",
			fields{client: testClient, database: database},
			args{input: Blogs.UserModel{
				Id:        userIdList[1],
				Role:      "admin",
				FirstName: "Bekhzod",
				LastName:  "Khudoyarov",
				Username:  "beccoder",
				Password:  Blogs.GeneratePasswordHash("qwerty"),
				Email:     "example@gmail.com",
				Bio:       "I am admin",
			}},
			primitive.ObjectID{},
			true,
		},
		{
			"CreateUserTestAdminSuccess2",
			fields{client: testClient, database: database},
			args{input: Blogs.UserModel{
				Id:        userIdList[2],
				Role:      "admin",
				FirstName: "Bekhzod",
				LastName:  "Khudoyarov",
				Username:  "example",
				Password:  Blogs.GeneratePasswordHash("qwerty"),
				Email:     "example@gmail.com",
				Bio:       "I am admin",
			}},
			userIdList[2],
			false,
		},
		{
			"CreateUserTestUserSuccess3",
			fields{client: testClient, database: database},
			args{input: Blogs.UserModel{
				Id:        userIdList[3],
				Role:      "user",
				FirstName: "Userbek",
				LastName:  "Userov",
				Username:  "user",
				Password:  Blogs.GeneratePasswordHash("user"),
				Email:     "user@gmail.com",
				Bio:       "I am user",
			}},
			userIdList[3],
			false,
		},
		{
			"CreateUserTestAuthorSuccess4",
			fields{client: testClient, database: database},
			args{input: Blogs.UserModel{
				Id:        userIdList[4],
				Role:      "author",
				FirstName: "Authorbek",
				LastName:  "Authorov",
				Username:  "author",
				Password:  Blogs.GeneratePasswordHash("author"),
				Email:     "author@gmail.com",
				Bio:       "I am author",
			}},
			userIdList[4],
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdmRepo{
				client:   tt.fields.client,
				database: tt.fields.database,
			}
			got, err := r.CreateUser(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdmRepo_GetAllUsers(t *testing.T) {
	type fields struct {
		client   *mongo.Client
		database string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			"GetAllUsersSuccess",
			fields{client: testClient, database: database},
			4,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdmRepo{
				client:   tt.fields.client,
				database: tt.fields.database,
			}
			response, err := r.GetAllUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(response) != tt.want {
				t.Errorf("GetAllUsers() got = %v, want %v", len(response), tt.want)
			}
		})
	}
}

func TestAdmRepo_GetUserById(t *testing.T) {
	type fields struct {
		client   *mongo.Client
		database string
	}
	type args struct {
		userId primitive.ObjectID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Blogs.UserResponse
		wantErr bool
	}{
		{
			"GetUserById_1",
			fields{client: testClient, database: database},
			args{userId: userIdList[0]},
			Blogs.UserResponse{
				Id:        userIdList[0],
				Role:      "admin",
				FirstName: "Bekhzod",
				LastName:  "Khudoyarov",
				Username:  "beccoder",
				Email:     "bekhzodkhudoyarov@gmail.com",
				Bio:       "I am admin",
			},
			false,
		},
		{
			"GetUserById_2",
			fields{client: testClient, database: database},
			args{userId: userIdList[2]},
			Blogs.UserResponse{
				Id:        userIdList[2],
				Role:      "admin",
				FirstName: "Bekhzod",
				LastName:  "Khudoyarov",
				Username:  "example",
				Email:     "example@gmail.com",
				Bio:       "I am admin",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdmRepo{
				client:   tt.fields.client,
				database: tt.fields.database,
			}
			got, err := r.GetUserById(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.CreatedAt = got.CreatedAt
			tt.want.UpdatedAt = got.UpdatedAt
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdmRepo_UpdateUser(t *testing.T) {
	type fields struct {
		client   *mongo.Client
		database string
	}
	type args struct {
		userId primitive.ObjectID
		input  Blogs.UpdateUserRequest
	}
	UpdatedStr := "My updated firstname"
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"UpdateUserSuccess_1",
			fields{client: testClient, database: database},
			args{
				userId: userIdList[0],
				input:  Blogs.UpdateUserRequest{FirstName: &UpdatedStr},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdmRepo{
				client:   tt.fields.client,
				database: tt.fields.database,
			}
			if err := r.UpdateUser(tt.args.userId, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdmRepo_DeleteUser(t *testing.T) {
	type fields struct {
		client   *mongo.Client
		database string
	}
	type args struct {
		userId primitive.ObjectID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"DeleteUserSuccess_1",
			fields{client: testClient, database: database},
			args{userId: userIdList[2]},
			false,
		},
		{
			"DeleteUserFail_1",
			fields{client: testClient, database: database},
			args{userId: userIdList[2]},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdmRepo{
				client:   tt.fields.client,
				database: tt.fields.database,
			}
			if err := r.DeleteUser(tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdmRepo_GetUserById_2(t *testing.T) {
	type fields struct {
		client   *mongo.Client
		database string
	}
	type args struct {
		userId primitive.ObjectID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Blogs.UserResponse
		wantErr bool
	}{
		{
			"GetUserById_3",
			fields{client: testClient, database: database},
			args{userId: userIdList[0]},
			Blogs.UserResponse{
				Id:        userIdList[0],
				Role:      "admin",
				FirstName: "My updated firstname",
				LastName:  "Khudoyarov",
				Username:  "beccoder",
				Email:     "bekhzodkhudoyarov@gmail.com",
				Bio:       "I am admin",
			},
			false,
		},
		{
			"GetUserByIdFail_4",
			fields{client: testClient, database: database},
			args{userId: userIdList[2]},
			Blogs.UserResponse{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdmRepo{
				client:   tt.fields.client,
				database: tt.fields.database,
			}
			got, err := r.GetUserById(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.CreatedAt = got.CreatedAt
			tt.want.UpdatedAt = got.UpdatedAt
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
