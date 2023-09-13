package main

import (
	"fiber_postgres/database"
	"fiber_postgres/model"
	"fiber_postgres/router"
	"fiber_postgres/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := server.Create()

	database.DB.AutoMigrate(&model.Book{})

	router.AppSetup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
