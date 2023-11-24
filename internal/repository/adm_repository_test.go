package repository

import (
	"Blogs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

var (
	idList = []primitive.ObjectID{primitive.NewObjectID(), primitive.NewObjectID(), primitive.NewObjectID()}
	testDB = StartTest()
)

func TestAdmRepo_CreateUser(t *testing.T) {
	type fields struct {
		db *mongo.Client
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
			fields{db: testDB},
			args{input: Blogs.UserModel{
				Id:        idList[0],
				Role:      "admin",
				FirstName: "Bekhzod",
				LastName:  "Khudoyarov",
				Username:  "beccoder",
				Password:  Blogs.GeneratePasswordHash("qwerty"),
				Email:     "bekhzodkhudoyarov@gmail.com",
				Bio:       "I am admin",
			}},
			idList[0],
			false,
		},
		{
			"CreateUserTestAdminFailDuplicateUsername",
			fields{db: testDB},
			args{input: Blogs.UserModel{
				Id:        idList[1],
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
			fields{db: testDB},
			args{input: Blogs.UserModel{
				Id:        idList[2],
				Role:      "admin",
				FirstName: "Bekhzod",
				LastName:  "Khudoyarov",
				Username:  "example",
				Password:  Blogs.GeneratePasswordHash("qwerty"),
				Email:     "example@gmail.com",
				Bio:       "I am admin",
			}},
			idList[2],
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdmRepo{
				db: tt.fields.db,
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
		db *mongo.Client
	}
	tests := []struct {
		name   string
		fields fields
		//want    []Blogs.UserResponse
		want    int
		wantErr bool
	}{
		{
			"GetAllUsersSuccess",
			fields{db: testDB},
			2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdmRepo{
				db: tt.fields.db,
			}
			response, err := r.GetAllUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(response) != 2 {
				t.Errorf("GetAllUsers() got = %v, want %v", len(response), tt.want)
			}
		})
	}
}

func TestAdmRepo_GetUserById(t *testing.T) {
	type fields struct {
		db *mongo.Client
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
			fields{db: testDB},
			args{userId: idList[0]},
			Blogs.UserResponse{
				Id:        idList[0],
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
			fields{db: testDB},
			args{userId: idList[2]},
			Blogs.UserResponse{
				Id:        idList[2],
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
				db: tt.fields.db,
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
		db *mongo.Client
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
			fields{db: testDB},
			args{
				userId: idList[0],
				input:  Blogs.UpdateUserRequest{FirstName: &UpdatedStr},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdmRepo{
				db: tt.fields.db,
			}
			if err := r.UpdateUser(tt.args.userId, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdmRepo_DeleteUser(t *testing.T) {
	type fields struct {
		db *mongo.Client
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
			fields{db: testDB},
			args{userId: idList[2]},
			false,
		},
		{
			"DeleteUserFail_1",
			fields{db: testDB},
			args{userId: idList[2]},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdmRepo{
				db: tt.fields.db,
			}
			if err := r.DeleteUser(tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdmRepo_GetUserById_2(t *testing.T) {
	type fields struct {
		db *mongo.Client
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
			fields{db: testDB},
			args{userId: idList[0]},
			Blogs.UserResponse{
				Id:        idList[0],
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
			fields{db: testDB},
			args{userId: idList[2]},
			Blogs.UserResponse{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdmRepo{
				db: tt.fields.db,
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
