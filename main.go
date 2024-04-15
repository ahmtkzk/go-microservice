package main

import (
	"go-microservice/internal/database"
	"go-microservice/internal/server"
)

func main() {
	db := database.NewDatabaseConnection()
	pr := database.NewProductRepository(db)
	sv := server.NewServerConnection(pr)
	sv.RegisterRoutes()
	sv.StartServer()
}
