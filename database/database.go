package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MyMongoDB struct {
	client            *mongo.Client
	context           context.Context
	cancel_function   context.CancelFunc
	safe_mode         bool
	_danger_databases []string
}

func MongoDB(host, port, user, password string) (*MyMongoDB, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", user, password, host, port) // make a python like format function

	the_context, the_cancel_function := context.WithCancel(context.Background())

	client, err := mongo.Connect(the_context, options.Client().ApplyURI(uri))

	_danger_databases := [...]string{"admin", "config", "local"}
	myMongoDB := MyMongoDB{client: client, context: the_context, cancel_function: the_cancel_function, safe_mode: true, _danger_databases: _danger_databases[:]}

	return &myMongoDB, err
}

func (self *MyMongoDB) Stop() {
	defer self.cancel_function()
	if err := self.client.Disconnect(self.context); err != nil {
		panic(err)
	}
}

func (self *MyMongoDB) Get_database(database_name string) (*mongo.Database, error) {
	database := self.client.Database(database_name)
	return database, nil
}

func (self *MyMongoDB) List_database() ([]string, error) {
	list, err := self.client.ListDatabaseNames(self.context, bson.M{})
	return list, err
}

func (self *MyMongoDB) Detele_database(database_name string) error {
	if self.safe_mode {
		if _, Found := find_string_in_a_list(self._danger_databases, database_name); Found == false {
			return self.client.Database(database_name).Drop(self.context)
		}
	} else {
		return self.client.Database(database_name).Drop(self.context)
	}
	return nil
}

func find_string_in_a_list(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
