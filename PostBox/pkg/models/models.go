package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

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
	Password   string
	Created_on time.Time
	Active     bool
}
