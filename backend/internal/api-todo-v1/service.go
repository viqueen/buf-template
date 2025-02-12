package api_todo_v1

import (
	v1 "api/go-sdk/todo/v1"
	"api/go-sdk/todo/v1/todoV1connect"
	"backend/internal/store/gendb"
	connect "connectrpc.com/connect"
	"context"
	"errors"
	"github.com/gofrs/uuid"
)

type todoService struct {
	store gendb.Querier
}

func NewTodoService(store gendb.Querier) todoV1connect.TodoServiceHandler {
	return &todoService{
		store: store,
	}
}

func (t todoService) CreateTodo(ctx context.Context, c *connect.Request[v1.CreateTodoRequest]) (*connect.Response[v1.CreateTodoResponse], error) {
	id := uuid.Must(uuid.NewV4())
	created, err := t.store.CreateTodo(ctx, &gendb.CreateTodoParams{
		ID:          id,
		Description: c.Msg.GetDescription(),
	})
	if err != nil {
		return nil, dbErrorToApi(err, "failed to create todo")
	}
	return connect.NewResponse(&v1.CreateTodoResponse{
		Todo: dbTodoToApi(created),
	}), nil

}

func (t todoService) GetTodo(ctx context.Context, c *connect.Request[v1.GetTodoRequest]) (*connect.Response[v1.GetTodoResponse], error) {
	id := uuid.FromStringOrNil(c.Msg.GetId())
	if id.IsNil() {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid id"))
	}
	found, err := t.store.GetTodo(ctx, id)
	if err != nil {
		return nil, dbErrorToApi(err, "failed to get todo")
	}
	return connect.NewResponse(&v1.GetTodoResponse{
		Todo: dbTodoToApi(found),
	}), nil
}

func (t todoService) ListTodos(ctx context.Context, c *connect.Request[v1.ListTodosRequest]) (*connect.Response[v1.ListTodosResponse], error) {
	todos, err := t.store.ListTodos(ctx, &gendb.ListTodosParams{
		PageLimit:  nonZeroOrDefaultInt32(c.Msg.GetPageLimit(), 10),
		PageOffset: c.Msg.GetPageOffset(),
	})
	if err != nil {
		return nil, dbErrorToApi(err, "failed to list todos")
	}
	var apiTodos []*v1.Todo
	for _, todo := range todos {
		apiTodos = append(apiTodos, dbTodoToApi(todo))
	}
	return connect.NewResponse(&v1.ListTodosResponse{
		Todos: apiTodos,
	}), nil
}
