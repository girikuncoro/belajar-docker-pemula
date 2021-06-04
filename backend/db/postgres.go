package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/girikuncoro/belajar-docker-pemula/backend/schema"
)

type Postgres struct {
	DB *sql.DB
}

func (p *Postgres) GetAll() ([]schema.Todo, error) {
	query := `
		SELECT *
		FROM todo
		ORDER BY id;
	`

	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}

	todoList := []schema.Todo{}
	for rows.Next() {
		var t schema.Todo
		if err := rows.Scan(&t.ID, &t.Note, &t.Done); err != nil {
			return nil, err
		}
		todoList = append(todoList, t)
	}
	return todoList, nil
}

func (p *Postgres) Insert(todo *schema.Todo) (int, error) {
	query := `
		INSERT INTO todo (id, note, done)
		VALUES(nextval('todo_id'), $1, $2)
		RETURNING id;
	`

	rows, err := p.DB.Query(query, todo.Note, convertBoolToBit(todo.Done))
	if err != nil {
		return -1, err
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return -1, err
		}
	}
	return id, nil
}

func (p *Postgres) Update(todo *schema.Todo) error {
	query := `
		UPDATE todo
		SET note = $2, done = $3
		WHERE id = $1;
	`

	rows, err := p.DB.Query(query, todo.ID, todo.Note, convertBoolToBit(todo.Done))
	if err != nil {
		return err
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return err
		}
	}
	return nil
}

func (p *Postgres) Delete(id int) error {
	query := `
		DELETE FROM todo
		WHERE id = $1;
	`

	if _, err := p.DB.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Close() {
	p.DB.Close()
}

func ConnectPostgres() (*Postgres, error) {
	connStr, err := loadPostgresConfig()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Postgres{db}, nil
}

func loadPostgresConfig() (string, error) {
	if os.Getenv("DB_HOST") == "" {
		return "", fmt.Errorf("Environment variable DB_HOST must be set")
	}
	if os.Getenv("DB_PORT") == "" {
		return "", fmt.Errorf("Environment variable DB_PORT must be set")
	}
	if os.Getenv("DB_USER") == "" {
		return "", fmt.Errorf("Environment variable DB_USER must be set")
	}
	if os.Getenv("DB_DATABASE") == "" {
		return "", fmt.Errorf("Environment variable DB_DATABASE must be set")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		return "", fmt.Errorf("Environment variable DB_PASSWORD must be set")
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	return connStr, nil
}

func convertBoolToBit(val bool) int {
	if val {
		return 1
	}
	return 0
}
