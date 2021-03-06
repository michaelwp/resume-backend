package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func DbCon(dbName string) (*mongo.Database, string){
	/*user := helpers.LoadEnv("DB_USER")
	pass := helpers.LoadEnv("DB_PASS")
	db := helpers.LoadEnv("DB_NAME")
	host := helpers.LoadEnv("DB_HOST")
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		user, pass, host, db)*/
	
	uri := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil { log.Fatal(err) }

	err = client.Ping(ctx, nil)
	if err != nil { log.Fatal(err) }

	status := "Connected to MongoDB!"

	return client.Database(dbName), status
}

