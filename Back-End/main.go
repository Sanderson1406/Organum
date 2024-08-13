package main

import (
	database "organum/DataBase"
	server "organum/Server"
)

func main() {
	database.ConectMongodb()
	server.StartServer()
}
