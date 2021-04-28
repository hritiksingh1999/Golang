package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func db() *mongo.Client {
	clientoptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, err := mongo.Connect(context.TODO(), clientoptions)
	if err != nil {
		log.Fatal(err)
	}
	err = ctx.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return ctx
}
