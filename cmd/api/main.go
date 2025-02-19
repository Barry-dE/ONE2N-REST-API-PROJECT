package main

import (
	"log"
	"time"

	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/db"
	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/env"
	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/handler"
	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/repository"
	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
)

func main() {
	// Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// App congiguration
	cfg := Config{
		addr: env.GetString("ADDR", ":3000"),
		env:  env.GetString("ENV", "development"),
		db: dbConfig{
			addr:               env.GetString("DB_ADDR", " "),
			maxOpenConnections: env.GetInt("DB_MAX_OPEN_CONNECTIONS", 20),
			maxIdleConnections: env.GetInt("DB_MAX_IDLE_CONNECTIONS", 10),
			maxIdleTime:        time.Minute * 10,
			MaxLifetime:        time.Hour,
		},
	}

	// Logger

	// Database
	db, err := db.NewDBConnection(cfg.db.addr, cfg.db.maxOpenConnections, cfg.db.maxIdleConnections, cfg.db.maxIdleTime, cfg.db.MaxLifetime)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("database connection established")

	repo := repository.NewStudentStore(db)

	h := handler.NewHandler(*repo)

	app := &application{
		config:  cfg,
		handler: h,
	}

	handler := app.mount(repo)

	if err := app.run(handler); err != nil {
		log.Fatal(err)
	}

	log.Printf("Server running on %s", cfg.addr)
}
