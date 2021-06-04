package db

import (
	"context"

	"github.com/girikuncoro/belajar-docker-pemula/backend/schema"
)

const repoKey = "repoKey"

type Repo interface {
	GetAll() ([]schema.Todo, error)
	Insert(todo *schema.Todo) (int, error)
	Update(todo *schema.Todo) error
	Delete(id int) error
	Close()
}

func SetRepo(ctx context.Context, repo Repo) context.Context {
	return context.WithValue(ctx, repoKey, repo)
}

func getRepo(ctx context.Context) Repo {
	return ctx.Value(repoKey).(Repo)
}

func GetAll(ctx context.Context) ([]schema.Todo, error) {
	return getRepo(ctx).GetAll()
}

func Insert(ctx context.Context, todo *schema.Todo) (int, error) {
	return getRepo(ctx).Insert(todo)
}

func Update(ctx context.Context, todo *schema.Todo) error {
	return getRepo(ctx).Update(todo)
}

func Delete(ctx context.Context, id int) error {
	return getRepo(ctx).Delete(id)
}

func Close(ctx context.Context) {
	getRepo(ctx).Close()
}
