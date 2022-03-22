package postgres

import (
	"context"
	"errors"

	"github.com/YoNoSoyVictor/ThoughtBox/pkg/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Signup(name, email, password string) error {
	//Inserts a user

	//create hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	statement := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3);"

	_, err = m.DB.Exec(context.Background(), statement, name, email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) Login(email, password string) (int, error) {

	var id int
	var hashedPassword []byte

	statement := "SELECT id, password FROM users WHERE email = $1 AND active = TRUE;"

	row := m.DB.QueryRow(context.Background(), statement, email)
	err := row.Scan(&id, &hashedPassword)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}
