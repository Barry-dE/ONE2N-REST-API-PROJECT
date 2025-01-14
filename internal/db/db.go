package db

import (
	"context"
	"database/sql"
	"time"
)

func NewDBConnection(addr string, maxOpenConns, maxIdleConns int, maxIdleTime time.Duration, maxLifeTime time.Duration) (*sql.DB, error) {

	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxIdleTime(maxIdleTime)
	db.SetConnMaxLifetime(maxLifeTime)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
