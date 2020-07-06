package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/rbonnat/blockchain-in-go/server"
)

func main() {
	var err error

	// Fetch environment variables
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")

	// Launch http server
	err = server.Run(context.TODO(), port)
	if err != nil {
		log.Fatal(err)
	}
}
