package main

import (
	"github.com/arthuramorim04/go-activity-api.git/database"
	"github.com/arthuramorim04/go-activity-api.git/server"
)

func main() {

	database.StartDB()
	server := server.NewServer()
	server.Run()

}
