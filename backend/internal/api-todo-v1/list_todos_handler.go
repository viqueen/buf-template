package api_todo_v1

import (
	"context"

	"connectrpc.com/connect"
	todoV1 "github.com/viqueen/buf-template/api/go-sdk/todo/v1"
	"github.com/viqueen/buf-template/backend/internal/store"
)

func (t todoService) ListTodos(
	ctx context.Context,
	request *connect.Request[todoV1.ListTodosRequest],
) (*connect.Response[todoV1.ListTodosResponse], error) {
	todos, err := t.repo.List(ctx, store.ListOptions{
		Limit:  nonZeroOrDefaultInt32(request.Msg.GetPageLimit(), PageLimitDefault),
		Offset: request.Msg.GetPageOffset(),
	})
	if err != nil {
		return nil, dbErrorToAPI(err, "failed to list todos")
	}

	apiTodos := make([]*todoV1.Todo, 0, len(todos))
	for index, todo := range todos {
		apiTodos[index] = dbTodoToAPI(todo)
	}

	return connect.NewResponse(&todoV1.ListTodosResponse{
		Todos: apiTodos,
	}), nil
}