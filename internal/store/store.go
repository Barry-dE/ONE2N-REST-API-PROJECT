package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type StudentRepository interface {
	Create(ctx context.Context, student *Student) error
}

type Storage struct {
	Students StudentRepository
}

func NewStudentStore(db *sql.DB) *Storage {
	return &Storage{
		Students: &StudentStore{db: db},
	}
}
