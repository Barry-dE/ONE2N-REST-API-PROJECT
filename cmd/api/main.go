package main

import (
	"log"

	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/env"
	"github.com/lpernett/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{
		addr: env.GetString("ADDR", ":3000"),
		env:  env.GetString("ENV", "Development"),
	}

	app := &application{
		config: cfg,
	}

	// Mount routes and get the gin engine
	handler := app.mount()

	//Start the server
	if err := app.run(handler); err != nil {
		log.Fatal(err)
	}
}
