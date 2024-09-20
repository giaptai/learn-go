package connectdb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnMongoDB(){
	// client, err:=mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/").SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)))
	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := Client.Database("db_comment").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

