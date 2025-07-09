package api_todo_v1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/gofrs/uuid"
	todoV1 "github.com/viqueen/buf-template/api/go-sdk/todo/v1"
	"github.com/viqueen/buf-template/backend/internal/store"
)

func (t todoService) CreateTodo(
	ctx context.Context,
	request *connect.Request[todoV1.CreateTodoRequest],
) (*connect.Response[todoV1.CreateTodoResponse], error) {
	id := uuid.Must(uuid.NewV4())

	todo := &store.Todo{
		ID:          id,
		Description: request.Msg.GetDescription(),
	}

	err := t.repo.Create(ctx, todo)
	if err != nil {
		return nil, dbErrorToAPI(err, "failed to create todo")
	}

	return connect.NewResponse(&todoV1.CreateTodoResponse{
		Todo: dbTodoToAPI(todo),
	}), nil
}