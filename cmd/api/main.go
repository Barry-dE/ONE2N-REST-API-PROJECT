package main

import (
	"log"

	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/env"
	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/store"
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

	// To-do: create the database connection pool and pass it to the store.
	
	store := store.NewStudentStore(nil)

	app := &application{
		config: cfg,
		store:  *store,
	}

	handler := app.mount()

	if err := app.run(handler); err != nil {
		log.Fatal(err)
	}
}
