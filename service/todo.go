package service

import (
	"context"

	"github.com/ThomasHamilton2/todo/db"
	"github.com/ThomasHamilton2/todo/schema"
)

func Close(ctx context.Context) {
	db.Close(ctx)
}

func Insert(ctx context.Context, todo *schema.Todo) (int, error) {
	return db.Insert(ctx, todo)
}

func Delete(ctx context.Context, id int) error {
	return db.Delete(ctx, id)
}

func GetAll(ctx context.Context) ([]schema.Todo, error) {
	return db.GetAll(ctx)
}
