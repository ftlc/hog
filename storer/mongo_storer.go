package storer

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// MongoStorer implements the Storer interface
// for a MongoDB back end
type MongoStorer struct {
	client   *mongo.Client
	database *mongo.Database
}

// Initialize does setup for MongoDB client
func (mongoStorer *MongoStorer) Initialize() error {
	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		return err
	}
	mongoStorer.client = client
	err = client.Connect(context.TODO())
	if err != nil {
		return err
	}
	mongoStorer.database = client.Database("hog")
	return nil
}

// SaveEntries saves a root-level array of JSON
// to MongoDB client under the "hog" database,
// with collection "tag"
func (mongoStorer *MongoStorer) SaveEntries(reader io.Reader, tag string) error {
	received, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatalln("Failed to read response from reader")
		return err
	}
	documents := make([]map[string]interface{}, 0)
	if err = json.Unmarshal(received, &documents); err != nil {
		log.Printf("Failed to unmarshal %s into %T", received, documents)
		log.Printf("Please ensure reader is a list of JSON")
		return err
	}
	return mongoStorer.castAndInsert(documents, tag)
}

func (mongoStorer *MongoStorer) castAndInsert(documents []map[string]interface{}, tag string) error {
	docsAsInterfaces := make([]interface{}, len(documents))
	for i, doc := range documents {
		docsAsInterfaces[i] = interface{}(doc)
	}
	collection := mongoStorer.database.Collection(tag)
	_, err := collection.InsertMany(context.Background(), docsAsInterfaces)
	return err
}
