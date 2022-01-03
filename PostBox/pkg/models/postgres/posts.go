package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/YoNoSoyVictor/ThoughtBox/pkg/models"
)

type PostModel struct {
	DB *pgxpool.Pool
}

func (m *PostModel) Insert(title, message string) (int, error) {
	//Inserts a Post (title and message, strings) into the database, returns id of newly created post

	statement := "INSERT INTO posts (title, message, created_on) VALUES ($1, $2, NOW()) RETURNING id;"
	var lastInsertId int

	err := m.DB.QueryRow(context.Background(), statement, title, message).Scan(&lastInsertId)
	if err != nil {
		return 0, nil
	}

	return lastInsertId, nil
}

func (m *PostModel) Get(id int) (*models.Post, error) {
	statement := "SELECT id, title, message, created_on FROM posts WHERE id = $1"

	row := m.DB.QueryRow(context.Background(), statement, id)

	p := &models.Post{}

	err := row.Scan(&p.ID, &p.Title, &p.Message, &p.Created_on)
	if err != nil {
		if err.Error() == "no rows in result set" {
			p.Message = "This post does not exist!"
		} else {
			return nil, err
		}
	}
	return p, nil
}

func (m *PostModel) Latest(ID int) ([]*models.Post, error) {
	statement := "SELECT id, title, message, created_on FROM posts ORDER BY id DESC LIMIT 10 OFFSET $1;"

	rows, err := m.DB.Query(context.Background(), statement, ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []*models.Post{}
	for rows.Next() {
		p := &models.Post{}
		err = rows.Scan(&p.ID, &p.Title, &p.Message, &p.Created_on)
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
