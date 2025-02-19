package repository

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
	ErrNotFound       = errors.New("resource not found")
)

type StudentRepository interface {
	Create(context.Context, *Student) error
	GetByID(context.Context, int64) (*Student, error)
	Update(context.Context, *Student) error
}

type Storage struct {
	Students StudentRepository
}

func NewStudentStore(db *sql.DB) *Storage {
	return &Storage{
		Students: &StudentStore{db: db},
	}
}
