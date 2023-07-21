package main

import (
	"bispy-agent/database"
	"bispy-agent/ticker"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[main]: Error loading .env file")
	}

	database.Connect()
	defer database.DB.Close()

	ticker.Start()
}
