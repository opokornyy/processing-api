package repository

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound  = errors.New("resource not found")
	ErrConfict   = errors.New("resource already exists")
	ErrEmptyData = errors.New("cannot store empty data")
	ErrInternal  = errors.New("database internal error")
)

type RepositoryError struct {
	RepositoryName string
	Message        string
	Err            error
}

func (e RepositoryError) Error() string {
	return fmt.Sprintf("%s: %s: %v", e.RepositoryName, e.Message, e.Err)
}

func (e RepositoryError) Unwrap() error {
	return e.Err
}
