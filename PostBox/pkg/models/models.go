package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Post struct {
	ID         int
	Title      string
	Message    string
	Created_on time.Time
}

type User struct {
	ID         int
	Name       string
	Email      string
	Password   []byte
	Created_on time.Time
	Active     bool
}
