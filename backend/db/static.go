package db

import (
	"github.com/girikuncoro/belajar-docker-pemula/backend/schema"
)

type Static struct{}

func (s *Static) GetAll() ([]schema.Todo, error) {
	todoList := []schema.Todo{
		{
			ID:   1,
			Note: "Beli mie instan",
			Done: false,
		},
		{
			ID:   2,
			Note: "Isi pulsa",
			Done: true,
		},
		{
			ID:   3,
			Note: "Ambil uang di atm",
			Done: false,
		},
	}
	return todoList, nil
}

func (s *Static) Insert(todo *schema.Todo) (int, error) {
	return 0, nil
}

func (s *Static) Update(todo *schema.Todo) error {
	return nil
}

func (s *Static) Delete(id int) error {
	return nil
}

func (s *Static) Close() {}
