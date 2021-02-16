package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection interface {
	Client() *mongo.Client
	DB() *mongo.Database
}

type conn struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewConnection() Connection {
	var c conn
	var err error

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(
		os.Getenv(("DB_URI")),
	))

	err = client.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		log.Fatal("Couldn't connect to the database ", err)
	} else {
		log.Println("MongoDb Connected!")
	}

	c.client = client
	c.db = client.Database(os.Getenv("DB_NAME"))

	return &c
}

func (c *conn) Client() *mongo.Client {
	return c.client
}


func (c *conn) DB() *mongo.Database {
	return c.db
}
