package mytests

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//https://www.loginradius.com/blog/async/mongodb-as-datasource-in-golang/

type User struct {
	Name     string
	Age      int
	Birthday time.Time
}

var uri = "mongodb://root:yingshaoxo666@127.0.0.1:27017/"
var client *mongo.Client
var err error
var cancel context.CancelFunc
var my_context context.Context
var collection *mongo.Collection

func TestMain(m *testing.M) {
	// os.Exit() does not respect defer statements

	client, err = mongo.Connect(my_context, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(my_context); err != nil {
			panic(err)
		}
	}()

	my_context, cancel = context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()

	collection = client.Database("testing").Collection("users")

	code := m.Run()
	fmt.Printf("the code of TestMain returns: %v", code)
	os.Exit(code)
}

func Test_datbase_connection(t *testing.T) {
	if err := client.Ping(my_context, readpref.Primary()); err != nil {
		t.Fatalf("can't connect to the mongodb service")
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
}

func Test_datbase_insertion(t *testing.T) {
	book1 := User{"yingshaoxo", 100, time.Now()}
	insertResult, err := collection.InsertOne(my_context, book1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func Test_database_list(t *testing.T) {
	filterCursor, _ := collection.Find(my_context, bson.M{})
	var episodesFiltered []bson.M
	if err = filterCursor.All(my_context, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesFiltered)
}

func Test_database_list_with_filter(t *testing.T) {
	filterCursor, _ := collection.Find(my_context, bson.M{"name": "yingshaoxo"})
	var episodesFiltered []bson.M
	if err = filterCursor.All(my_context, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesFiltered)
}

func Test_database_deletion(t *testing.T) {
	collection.DeleteMany(my_context, bson.M{})
}
