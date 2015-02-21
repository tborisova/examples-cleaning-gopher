package main

import (
	"fmt"
	"github.com/tborisova/clean_like_gopher"
	"gopkg.in/mgo.v2"
	"testing"
)

var mongoStartOptions = map[string]string{"host": "localhost", "dbName": "test", "port": "27017"}

func TestAdd25UsersToMongoDb(t *testing.T) {
	cleaner, _ := clean_like_gopher.NewCleaningGopher("mongo", mongoStartOptions)
	session, _ := mgo.Dial("localhost:27017")
	defer session.Close()

	AddUsersToMongoDB(25, session)

	collection := session.DB("test").C("people")
	count, _ := collection.Count()

	if count != 25 {
		fmt.Println(count)
		t.Errorf("Expected 25 users to be added!")
	}

	cleaner.Clean(nil)
	cleaner.Close()
}

func TestAdd5UsersToMongoDb(t *testing.T) {
	cleaner, _ := clean_like_gopher.NewCleaningGopher("mongo", mongoStartOptions)
	session, _ := mgo.Dial("localhost:27017")
	defer session.Close()

	AddUsersToMongoDB(5, session)

	collection := session.DB("test").C("people")
	count, _ := collection.Count()

	if count != 5 {
		fmt.Println(count)

		t.Errorf("Expected 25 users to be added!")
	}

	cleaner.Clean(nil)
	cleaner.Close()
}
