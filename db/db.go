package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	db *mongo.Client
}

func NewConn(uri string) (*Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	return &Database{db: client}, nil
}

func (d *Database) Close() {
	if err := d.db.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func (d *Database) GetDB() *mongo.Client {
	return d.db
}
