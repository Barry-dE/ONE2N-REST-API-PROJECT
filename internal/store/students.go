package store

import (
	"context"
	"database/sql"
	"time"
)

type Student struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Age       int       `json:"age" binding:"required"`
	Sex       string    `json:"sex" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StudentStore struct {
	db *sql.DB
}

func (s *StudentStore) Create(ctx context.Context, student *Student) error {
	query := `
	INSERT INTO students (firstname, lastname, email, age, sex) VALUES($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(ctx, query, student.FirstName, student.LastName, student.Email, student.Age, student.Sex).Scan(
		&student.ID,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
