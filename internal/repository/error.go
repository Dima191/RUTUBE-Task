package repository

import "errors"

var (
	ErrEmployeeNotFound      = errors.New("employee not found")
	ErrEmployeeAlreadyExists = errors.New("employee already exists")
	ErrDatabaseQuery         = errors.New("database query error")
	ErrNoSession             = errors.New("error no session")
	ErrAlreadySubscribed     = errors.New("already subscribed")
)
