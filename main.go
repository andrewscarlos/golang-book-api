package main

import (
	"github.com/andrewscarlos/golang-book-api/database"
	server "github.com/andrewscarlos/golang-book-api/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
