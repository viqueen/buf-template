package api_todo_v1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/gofrs/uuid"
	todoV1 "github.com/viqueen/buf-template/api/go-sdk/todo/v1"
)

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

	found, err := t.repo.GetByID(ctx, identifier)
	if err != nil {
		return nil, dbErrorToAPI(err, "failed to get todo")
	}

	return connect.NewResponse(&todoV1.GetTodoResponse{
		Todo: dbTodoToAPI(found),
	}), nil
}
