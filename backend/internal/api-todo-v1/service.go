package api_todo_v1

import (
	"github.com/viqueen/buf-template/api/go-sdk/todo/v1/todoV1connect"
	"github.com/viqueen/buf-template/backend/internal/store"
)

type todoService struct {
	repo *store.TodoRepository
}

func NewTodoService(repo *store.TodoRepository) todoV1connect.TodoServiceHandler {
	return &todoService{
		repo: repo,
	}
}
