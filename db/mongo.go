package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "go.mongodb.org/mongo-driver/bson"
)

var (
    client *mongo.Client
	billsCollection *mongo.Collection
	databaseName = "supermarket"
	collectionName = "bills"
    databaseURL = "mongodb://localhost:27017"
)


func ConnectMongoDB() error {
    // Create a MongoDB client
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(databaseURL))
	if err != nil {
		return err
	}

	// Connect to the MongoDB server
	err = client.Connect(context.Background())
	if err != nil {
		return err
	}

    // Ping the MongoDB server to check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

    log.Println("Connected to MongoDB!")


    // Get the database
	database := client.Database(databaseName)

	// Check if the collection exists
	collections, err := database.ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	// Check if the collection already exists in the list of collections
	collectionExists := false
	for _, collName := range collections {
		if collName == collectionName {
			collectionExists = true
			break
		}
	}

	// Create the collection if it doesn't exist
	if !collectionExists {
		// Create a new collection
		err = client.Database(databaseName).CreateCollection(context.Background(), collectionName)
		if err != nil {
			return err
		}
	}

    // Assign the billsCollection variable
    billsCollection = database.Collection(collectionName)

    return nil
}

func GetMongoClient() *mongo.Client {
    return client
}

func GetBillsCollection() *mongo.Collection {
	return billsCollection
}
