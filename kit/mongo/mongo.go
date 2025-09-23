package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client     *mongo.Client
	Collection *mongo.Collection
)

func ConnectMongo() (*mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := "mongodb+srv://yeremi:B4tmtd8B9zlyAgxW@easylife-domains-dev-cl.zhkc63h.mongodb.net/"
	clientOptions := options.Client().ApplyURI(uri)

	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error conectando a MongoDB: %w", err)
	}

	if err := Client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("error al hacer ping a MongoDB: %w", err)
	}

	db := Client.Database("easylife_customer_db")
	Collection = db.Collection("customer")

	return Collection, nil
}

func CloseMongo() {
	if Client != nil {
		if err := Client.Disconnect(context.Background()); err != nil {
			fmt.Printf("❌ Error cerrando MongoDB: %v\n", err)
		} else {
			fmt.Println("✅ Conexión a MongoDB cerrada")
		}
	}
}
