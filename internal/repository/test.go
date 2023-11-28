package repository

import (
	"Blogs"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

var (
	userIdList    = []primitive.ObjectID{primitive.NewObjectID(), primitive.NewObjectID(), primitive.NewObjectID(), primitive.NewObjectID(), primitive.NewObjectID()}
	postIdList    = []primitive.ObjectID{primitive.NewObjectID(), primitive.NewObjectID(), primitive.NewObjectID(), primitive.NewObjectID()}
	commentIdList = []primitive.ObjectID{primitive.NewObjectID(), primitive.NewObjectID(), primitive.NewObjectID(), primitive.NewObjectID()}

	posts = []Blogs.PostModel{
		{
			Title: "My title",
			Text:  "My text",
		},
		{
			Title: "My title 2",
		},
		{
			Text: "My text 3",
		},
		{
			Title: "My title of author",
			Text:  "My text of author",
		},
	}
	postsUpdates = []Blogs.PostModel{
		{
			Title: "My title update",
			Text:  "My text update",
		},
		{
			Title: "My title 2 update",
			Text:  "My text 2 updated ",
		},
		{
			Title: "My title 3 updated",
			Text:  "My text 3 update",
		},
	}
	comments = []Blogs.CommentModel{
		{
			Id:            commentIdList[0],
			PostId:        postIdList[0],
			CommentedById: userIdList[3],
			Comment:       "This is my first comment",
		},
		{
			Id:             commentIdList[1],
			PostId:         postIdList[0],
			CommentedById:  userIdList[4],
			ReplyCommentId: commentIdList[0],
			Comment:        "This is my reply comment for first comment",
		},
		{
			Id:             commentIdList[2],
			PostId:         postIdList[0],
			CommentedById:  userIdList[4],
			ReplyCommentId: commentIdList[0],
			Comment:        "This is my second reply comment for first comment",
		},
	}

	commentsUpdate = []string{
		"This is my first comment UPDATE",
		"This is my reply comment for first comment UPDATE",
		"This is my second reply comment for first comment UPDATE",
	}

	testClient, database = StartTest()
)

func StartTest() (*mongo.Client, string) {
	err := godotenv.Load(".env")
	if err != nil && !os.IsNotExist(err) {
		log.Println("failed to load .env file")
	}
	cfg := Blogs.Load()

	cfg.MakeMongoDBURL()
	client, err := ConnectMongoDB(cfg.MONGODB.URI, cfg.MONGODB.Username, cfg.MONGODB.Password)
	if err != nil {
		panic(err)
	}
	err = InitSchemas(client, cfg.MONGODB.Database)
	if err != nil {
		panic(err)
	}
	return client, cfg.MONGODB.Database
}
