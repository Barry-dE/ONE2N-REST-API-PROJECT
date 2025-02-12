package repository

import (
	"context"
	"database/sql"
	"log"

	"time"

	"github.com/lib/pq"
)

type Student struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Age       int       `json:"age" binding:"required,gte=0,lte=150"`
	Sex       string    `json:"sex" binding:"required,oneof=M F Other"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StudentStore struct {
	db *sql.DB
}

func (s *StudentStore) Create(ctx context.Context, student *Student) error {
	query := `
	INSERT INTO students ("firstName", "lastName", "email", "age", "sex") VALUES($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at
	`
	// use context to cancle long running mutations

	err := s.db.QueryRowContext(ctx, query, student.FirstName, student.LastName, student.Email, student.Age, student.Sex).Scan(
		&student.ID,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return ErrDuplicateEmail
			default:
				log.Printf("database error: %v (Code: %s)", pqErr.Message, pqErr.Code)
				return err
			}

		}
		return err
	}

	return nil
}
