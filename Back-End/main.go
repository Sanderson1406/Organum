package main

import (
	database "example.com/m/DataBase"
	server "example.com/m/Server"
)

func main() {
	database.ConectMongodb()
	server.StartServer()
}
