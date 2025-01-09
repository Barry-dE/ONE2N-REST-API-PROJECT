package main

import (
	"log"
	"time"

	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/db"
	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/env"
	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/store"
	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Application congiguration
	cfg := Config{
		addr: env.GetString("ADDR", ":3000"),
		env:  env.GetString("ENV", "Development"),
		db: dbConfig{
			addr:               env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/students?sslmode=disable"),
			maxOpenConnections: env.GetInt("DB_MAX_OPEN_CONNECTIONS", 20),
			maxIdleConnections: env.GetInt("DB_MAX_IDLE_CONNECTIONS", 10),
			maxIdleTime:        time.Minute * 10,
			MaxLifetime:        time.Hour,
		},
	}

	log.Printf("Attempting to start server on %s in %s environment...", cfg.addr, cfg.env)

	// Establsih a new database connection pool
	db, err := db.NewDBConnection(cfg.db.addr, cfg.db.maxOpenConnections, cfg.db.maxIdleConnections, cfg.db.maxIdleTime, cfg.db.MaxLifetime)
	if err != nil {
		log.Panic(err)
	}

	log.Println("database connection established")

	defer db.Close()

	// Initialize student data store
	store := store.NewStudentStore(db)

	app := &application{
		config: cfg,
		store:  *store,
	}

	handler := app.mount()

	if err := app.run(handler); err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started on %s", cfg.addr)
}
