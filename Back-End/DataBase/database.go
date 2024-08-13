package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConectMongodb() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb+srv://sandersonoficial10:EUamoBACKend@projeto.mb79irr.mongodb.net/?retryWrites=true&w=majority&appName=Projeto")
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado ao MongoDB!")
}

func GetMongoClient() *mongo.Client {
	clientInstance := client
	return clientInstance
}
