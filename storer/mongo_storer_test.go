package storer

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"testing"
)

func TestInitialize(t *testing.T) {
	mongoStorer := new(MongoStorer)
	err := mongoStorer.Initialize()
	if err != nil {
		t.Errorf("Could not connect to MongoDB: %s", err)
	}
}

func TestSaveEntriesMultiple(t *testing.T) {
	mongoStorer := new(MongoStorer)
	mongoStorer.Initialize()
	data := []byte(`[{"test":123},{"test":456}]`)
	reader := bytes.NewReader(data)
	err := mongoStorer.SaveEntries(reader, "sometag")
	if err != nil {
		t.Errorf("Could not store multiple entries: %s", err)
	}
	collection := mongoStorer.database.Collection("sometag")
	count, err := collection.Count(context.Background(), nil)
	if err != nil {
		t.Errorf("Could not get count")
	}
	if count < 2 {
		t.Errorf("No errors in implementation but no change in database")
	}
	collection.Drop(context.Background())
}

func TestSaveEntriesSingular(t *testing.T) {
	mongoStorer := new(MongoStorer)
	mongoStorer.Initialize()
	data := []byte(`{"test":123}`)
	reader := bytes.NewReader(data)
	log.SetOutput(ioutil.Discard) // disregard error logging for test
	err := mongoStorer.SaveEntries(reader, "sometag")
	if err == nil {
		t.Errorf("Should have failed to unmarshal data: %s", data)
	}
}
