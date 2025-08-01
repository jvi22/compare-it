package persistence

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo(uri string) error {
    var err error
    MongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
    return err
}

func SessionsCollection() *mongo.Collection {
    return MongoClient.Database("grocery").Collection("sessions")
}
