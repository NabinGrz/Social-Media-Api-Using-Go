package dbConnection

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://nabin:nabin@socialmedia.jicw8en.mongodb.net/?retryWrites=true&w=majority&appName=SocialMedia"
const dbName = "SocialMedia"
const userCollectionName = "user"
const postCollectionName = "post"

var UserCollection, PostCollection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal("🛑🛑🛑", err)
	}

	UserCollection = client.Database(dbName).Collection(userCollectionName)
	PostCollection = client.Database(dbName).Collection(postCollectionName)

	fmt.Println("🟢🟢🟢 Succesfully connected to MongoDB.....")
}
