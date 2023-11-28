package repository

import (
	"Blogs"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestAuthRepo_GetUser(t *testing.T) {
	type fields struct {
		client   *mongo.Client
		database string
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Blogs.UserResponse
		wantErr bool
	}{
		{
			"GetUserSuccess",
			fields{client: testClient, database: database},
			args{
				username: "beccoder",
				password: Blogs.GeneratePasswordHash("qwerty"),
			},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AuthRepo{
				client:   tt.fields.client,
				database: tt.fields.database,
			}
			got, err := r.GetUser(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.CreatedAt = got.CreatedAt
			tt.want.UpdatedAt = got.UpdatedAt
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
