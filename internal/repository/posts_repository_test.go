package repository

import (
	"Blogs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestPostsRepo_CreatePosts(t *testing.T) {
	type fields struct {
		db *mongo.Client
	}
	type args struct {
		post Blogs.PostModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    primitive.ObjectID
		wantErr bool
	}{
		{
			"CreatePostsTestSuccess_1",
			fields{db: testDB},
			args{post: Blogs.PostModel{
				Id:        postIdList[0],
				AuthorsId: userIdList[0],
				Title:     posts[0].Title,
				Text:      posts[0].Text,
			}},
			postIdList[0],
			false,
		},
		{
			"CreatePostsTestSuccess_2",
			fields{db: testDB},
			args{post: Blogs.PostModel{
				Id:        postIdList[1],
				AuthorsId: userIdList[0],
				Title:     posts[1].Title,
				Text:      posts[1].Text,
			}},
			postIdList[1],
			false,
		},
		{
			"CreatePostsTestSuccess_3",
			fields{db: testDB},
			args{post: Blogs.PostModel{
				Id:        postIdList[2],
				AuthorsId: userIdList[0],
				Title:     posts[2].Title,
				Text:      posts[2].Text,
			}},
			postIdList[2],
			false,
		},
		//{
		//	"CreatePostsTestFail_4",
		//	fields{db: testDB},
		//	args{post: Blogs.PostModel{
		//		Id:        postIdList[3],
		//		AuthorsId: userIdList[3],
		//		Title:     posts[2].Title,
		//		Text:      posts[2].Text,
		//	}},
		//	postIdList[3],
		//	true,
		//},
		{
			"CreatePostsTestSuccess_5",
			fields{db: testDB},
			args{post: Blogs.PostModel{
				Id:        postIdList[3],
				AuthorsId: userIdList[4],
				Title:     posts[3].Title,
				Text:      posts[3].Text,
			}},
			postIdList[3],
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostsRepo{
				db: tt.fields.db,
			}
			got, err := p.CreatePosts(tt.args.post)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePosts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostsRepo_GetMyAllPosts(t *testing.T) {
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
		want    []Blogs.PostResponse
		wantErr bool
	}{
		{
			"GetMyAllPosts_Success_1",
			fields{db: testDB},
			args{userId: userIdList[0]},
			[]Blogs.PostResponse{
				{
					Id:        postIdList[0],
					AuthorsId: userIdList[0],
					Title:     posts[0].Title,
					Text:      posts[0].Text,
					Likes:     nil,
				},
				{
					Id:        postIdList[1],
					AuthorsId: userIdList[0],
					Title:     posts[1].Title,
					Text:      posts[1].Text,
					Likes:     nil,
				},
				{
					Id:        postIdList[2],
					AuthorsId: userIdList[0],
					Title:     posts[2].Title,
					Text:      posts[2].Text,
					Likes:     nil,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostsRepo{
				db: tt.fields.db,
			}
			got, err := p.GetMyAllPosts(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMyAllPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want[0].CreatedAt = got[0].CreatedAt
			tt.want[1].CreatedAt = got[1].CreatedAt
			tt.want[2].CreatedAt = got[2].CreatedAt
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMyAllPosts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostsRepo_GetAllPosts(t *testing.T) {
	type fields struct {
		db *mongo.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Blogs.PostResponse
		wantErr bool
	}{
		{
			"GetAllPosts_Success_1",
			fields{db: testDB},
			[]Blogs.PostResponse{
				{
					Id:        postIdList[0],
					AuthorsId: userIdList[0],
					Title:     posts[0].Title,
					Text:      posts[0].Text,
					Likes:     nil,
				},
				{
					Id:        postIdList[1],
					AuthorsId: userIdList[0],
					Title:     posts[1].Title,
					Text:      posts[1].Text,
					Likes:     nil,
				},
				{
					Id:        postIdList[2],
					AuthorsId: userIdList[0],
					Title:     posts[2].Title,
					Text:      posts[2].Text,
					Likes:     nil,
				},
				{
					Id:        postIdList[3],
					AuthorsId: userIdList[4],
					Title:     posts[3].Title,
					Text:      posts[3].Text,
					Likes:     nil,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostsRepo{
				db: tt.fields.db,
			}
			got, err := p.GetAllPosts()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want[0].CreatedAt = got[0].CreatedAt
			tt.want[1].CreatedAt = got[1].CreatedAt
			tt.want[2].CreatedAt = got[2].CreatedAt
			tt.want[3].CreatedAt = got[3].CreatedAt
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllPosts() got:\n %v, want: \n %v", got, tt.want)
			}
		})
	}
}

func TestPostsRepo_GetPostById(t *testing.T) {
	type fields struct {
		db *mongo.Client
	}
	type args struct {
		postId primitive.ObjectID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Blogs.PostResponse
		wantErr bool
	}{
		{
			"GetPostById_Success_1",
			fields{db: testDB},
			args{postId: postIdList[0]},
			Blogs.PostResponse{
				Id:        postIdList[0],
				AuthorsId: userIdList[0],
				Title:     posts[0].Title,
				Text:      posts[0].Text,
				Likes:     nil,
			},
			false,
		},
		{
			"GetPostById_Success_2",
			fields{db: testDB},
			args{postId: postIdList[1]},
			Blogs.PostResponse{
				Id:        postIdList[1],
				AuthorsId: userIdList[0],
				Title:     posts[1].Title,
				Text:      posts[1].Text,
				Likes:     nil,
			},
			false,
		},
		{
			"GetPostById_Success_3",
			fields{db: testDB},
			args{postId: postIdList[2]},
			Blogs.PostResponse{
				Id:        postIdList[2],
				AuthorsId: userIdList[0],
				Title:     posts[2].Title,
				Text:      posts[2].Text,
				Likes:     nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostsRepo{
				db: tt.fields.db,
			}
			got, err := p.GetPostById(tt.args.postId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPostById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.CreatedAt = got.CreatedAt
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPostById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostsRepo_UpdatePost(t *testing.T) {
	type fields struct {
		db *mongo.Client
	}
	type args struct {
		postId primitive.ObjectID
		input  Blogs.UpdatePostRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"UpdatePost_Success_1",
			fields{db: testDB},
			args{
				postId: postIdList[0],
				input: Blogs.UpdatePostRequest{
					Title: &postsUpdates[0].Title,
					Text:  &postsUpdates[0].Text,
				},
			},
			false,
		},
		{
			"UpdatePost_Success_2",
			fields{db: testDB},
			args{
				postId: postIdList[1],
				input: Blogs.UpdatePostRequest{
					Title: &postsUpdates[1].Title,
					Text:  &postsUpdates[1].Text,
				},
			},
			false,
		},
		{
			"UpdatePost_Success_3",
			fields{db: testDB},
			args{
				postId: postIdList[2],
				input: Blogs.UpdatePostRequest{
					Title: &postsUpdates[2].Title,
					Text:  &postsUpdates[2].Text,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostsRepo{
				db: tt.fields.db,
			}
			if err := p.UpdatePost(tt.args.postId, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostsRepo_DeletePost(t *testing.T) {
	type fields struct {
		db *mongo.Client
	}
	type args struct {
		postId primitive.ObjectID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"DeletePost_Success_1",
			fields{db: testDB},
			args{postId: postIdList[2]},
			false,
		},
		{
			"DeletePost_Fail_2",
			fields{db: testDB},
			args{postId: postIdList[2]},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostsRepo{
				db: tt.fields.db,
			}
			if err := p.DeletePost(tt.args.postId); (err != nil) != tt.wantErr {
				t.Errorf("DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostsRepo_CreateComment(t *testing.T) {
	type fields struct {
		db *mongo.Client
	}
	type args struct {
		input Blogs.CommentModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    primitive.ObjectID
		wantErr bool
	}{
		{
			"CreateComment_Success_1",
			fields{db: testDB},
			args{input: Blogs.CommentModel{
				Id:            commentIdList[0],
				PostId:        postIdList[0],
				CommentedById: userIdList[3],
				Comment:       "This is my first comment",
			}},
			commentIdList[0],
			false,
		},
		{
			"CreateComment_Success_2",
			fields{db: testDB},
			args{input: Blogs.CommentModel{
				Id:             commentIdList[1],
				PostId:         postIdList[0],
				CommentedById:  userIdList[4],
				ReplyCommentId: commentIdList[0],
				Comment:        "This is my reply comment for first comment",
			}},
			commentIdList[1],
			false,
		},
		{
			"CreateComment_Success_3",
			fields{db: testDB},
			args{input: Blogs.CommentModel{
				Id:             commentIdList[2],
				PostId:         postIdList[0],
				CommentedById:  userIdList[4],
				ReplyCommentId: commentIdList[0],
				Comment:        "This is my second reply comment for first comment",
			}},
			commentIdList[2],
			false,
		},
		{
			"CreateComment_Fail",
			fields{db: testDB},
			args{input: Blogs.CommentModel{
				Id:             commentIdList[2],
				PostId:         postIdList[0],    // no post
				CommentedById:  userIdList[4],    // no user
				ReplyCommentId: commentIdList[0], // no reply comment
				Comment:        "Waiting error",
			}},
			primitive.ObjectID{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostsRepo{
				db: tt.fields.db,
			}
			got, err := p.CreateComment(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateComment() got = %v, want %v", got, tt.want)
			}
		})
	}
}
