package main

import (
	"github.com/bagashiz/freepass-2023/db"
	"github.com/bagashiz/freepass-2023/server"
)

func main() {
	// initiate database
	db.Connect()
	db.Migrate()

	// initiate server
	server.SetupRouter()
	server.Start(":8080")
}
