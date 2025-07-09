package api_todo_v1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/gofrs/uuid"
	todoV1 "github.com/viqueen/buf-template/api/go-sdk/todo/v1"
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

func (t todoService) CreateTodo(
	ctx context.Context,
	request *connect.Request[todoV1.CreateTodoRequest],
) (*connect.Response[todoV1.CreateTodoResponse], error) {
	id := uuid.Must(uuid.NewV4())

	todo := &store.Todo{
		ID:          id,
		Description: request.Msg.GetDescription(),
	}

	err := t.repo.CreateTodo(todo)
	if err != nil {
		return nil, dbErrorToAPI(err, "failed to create todo")
	}

	return connect.NewResponse(&todoV1.CreateTodoResponse{
		Todo: dbTodoToAPI(todo),
	}), nil
}

func (t todoService) GetTodo(
	ctx context.Context,
	request *connect.Request[todoV1.GetTodoRequest],
) (*connect.Response[todoV1.GetTodoResponse], error) {
	identifier := uuid.FromStringOrNil(request.Msg.GetId())
	if identifier.IsNil() {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			ErrInvalidTodoID,
		)
	}

	found, err := t.repo.GetTodo(identifier)
	if err != nil {
		return nil, dbErrorToAPI(err, "failed to get todo")
	}

	return connect.NewResponse(&todoV1.GetTodoResponse{
		Todo: dbTodoToAPI(found),
	}), nil
}

func (t todoService) ListTodos(
	ctx context.Context,
	request *connect.Request[todoV1.ListTodosRequest],
) (*connect.Response[todoV1.ListTodosResponse], error) {
	todos, err := t.repo.ListTodos(
		nonZeroOrDefaultInt32(request.Msg.GetPageLimit(), PageLimitDefault),
		request.Msg.GetPageOffset(),
	)
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
