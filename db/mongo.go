package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

var (
    client *mongo.Client
	billsCollection *mongo.Collection
	databaseName    = "supermarket"
	collectionName = "bills"
)

func ConnectMongoDB() error {
    clientOptions := options.Client().ApplyURI("mongodb+srv://username:password@cluster_name")
    ctx := context.TODO()

    var err error
    client, err = mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Ping the MongoDB server to check if the connection was successful
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")
	billsCollection = client.Database(databaseName).Collection(collectionName)

    return nil
}

func GetMongoClient() *mongo.Client {
    return client
}

func GetBillsCollection() *mongo.Collection {
	return billsCollection
}
