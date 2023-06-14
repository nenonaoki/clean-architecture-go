package main

import (
	sample_database "github.com/nenonaoki/clean-architecture-go/pkg/frameworks-drivers/database"
	sample_server "github.com/nenonaoki/clean-architecture-go/pkg/frameworks-drivers/server"
)

func main() {
	database := sample_database.NewSampleDatabase()
	app := sample_server.NewSampleServer(*database)
	app.Start()
}
