package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MI MongoInstance

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func (m *MongoInstance) Collection(name string) *mongo.Collection {
	return m.DB.Collection(name)
}

func ConnectMongoDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	MI = MongoInstance{
		Client: client,
		DB:     client.Database(os.Getenv("MONGO_DB_NAME")),
	}
}
