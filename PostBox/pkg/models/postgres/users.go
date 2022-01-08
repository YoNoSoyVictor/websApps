package postgres

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Signup(name, email, password string) (pgconn.CommandTag, error) {
	//Inserts a user

	statement := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3);"

	tag, err := m.DB.Exec(context.Background(), statement, name, email, password)
	if err != nil {
		return nil, err
	}

	return tag, nil
}
