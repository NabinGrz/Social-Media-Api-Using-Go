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
const collectionName = "user"

var SocialMediaCollection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal("🛑🛑🛑", err)
	}

	SocialMediaCollection = client.Database(dbName).Collection(collectionName)

	fmt.Println("🟢🟢🟢 Succesfully connected to MongoDB.....")
}
