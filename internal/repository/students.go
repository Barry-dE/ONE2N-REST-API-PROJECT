package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"time"

	util "github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/utils"
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
	Version   int       `json:"version"`
}

type StudentStore struct {
	db *sql.DB
}

func (s *StudentStore) Create(ctx context.Context, student *Student) error {
	query := `
	INSERT INTO students ("firstName", "lastName", email, age, sex) VALUES($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at
	`

	ctx, cancel := util.TimeoutCtx(ctx)
	defer cancel()

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

func (s *StudentStore) GetByID(ctx context.Context, studentID int64) (*Student, error) {
	query := `SELECT id, "firstName", "lastName", email, age, sex, created_at, updated_at, version
 FROM students
 WHERE id = $1`

	ctx, cancel := util.TimeoutCtx(ctx)
	defer cancel()

	student := &Student{}

	err := s.db.QueryRowContext(ctx, query, studentID).Scan(
		&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Age, &student.Sex, &student.CreatedAt, &student.UpdatedAt, &student.Version)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return student, nil

}

func (s *StudentStore) Update(ctx context.Context, student *Student) error {
	query := `
	UPDATE students
	SET "firstName" = $1, "lastName" = $2, email = $3, age = $4, sex = $5, version = version + 1
	WHERE id = $6 AND version = $7
	RETURNING version
	`
	ctx, cancel := util.TimeoutCtx(ctx)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, &student.FirstName, &student.LastName, &student.Email, &student.Age, &student.Sex, &student.ID, &student.Version).Scan(&student.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *StudentStore) Delete(ctx context.Context, studentID int64) error {
	query := `DELETE FROM  students WHERE id=$1`

	ctx, cancel := util.TimeoutCtx(ctx)
	defer cancel()

	result, err := s.db.ExecContext(ctx, query, studentID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}
