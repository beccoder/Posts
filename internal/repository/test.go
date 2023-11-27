package repository

import (
	"Blogs"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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

	testDB = StartTest()
)

func StartTest() *mongo.Client {
	if err := Blogs.LoadEnvConfig(); err != nil {
		log.Fatal(err)
	}

	dbURI := viper.GetString("MONGO.PROTOCOL") + "://" + viper.GetString("MONGO.HOST") + ":" + viper.GetString("MONGO.PORT")
	client, err := ConnectMongoDB(dbURI)
	if err != nil {
		panic(err)
	}
	err = InitSchemas(client)
	if err != nil {
		panic(err)
	}
	return client
}
