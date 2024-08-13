package main

import (
  "context"
  "fmt"
  //"go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
)

func ConectMongodb() {
  var err error
	clientOptions := options.Client().ApplyURI("mongodb+srv://sandersonoficial10:EUamoBACKend@projeto.mb79irr.mongodb.net/?retryWrites=true&w=majority&appName=Projeto")
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verificar a conexão
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado ao MongoDB!")
}
