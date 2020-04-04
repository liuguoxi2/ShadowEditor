package helper

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestMongo(t *testing.T) {
	config, err := GetConfig("../config.toml")
	if err != nil {
		t.Error(err)
		return
	}

	db, err := NewMongo(config.Database.Connection, "Test")
	if err != nil {
		t.Error(err)
		return
	}
	collectionName := "PersonTest"
	persons := []Person{
		{"xiaoming", 10},
		{"xiaoli", 20},
	}

	// insert
	_, err = db.InsertMany(collectionName, persons)
	if err != nil {
		t.Error(err)
		return
	}

	// find
	results := []Person{}
	err = db.FindMany(collectionName, bson.M{}, &results)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(results)

	// listCollectionNames
	collectionNames, err := db.ListCollectionNames()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(collectionNames)

	// drop
	err = db.DropCollection(collectionName)
	if err != nil {
		t.Error(err)
	}
}

type Person struct {
	name string
	age  int
}