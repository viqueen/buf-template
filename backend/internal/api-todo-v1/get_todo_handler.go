package api_todo_v1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
	todoV1 "github.com/viqueen/buf-template/api/go-sdk/todo/v1"
)

func (t todoService) GetTodo(
	ctx context.Context,
	request *connect.Request[todoV1.GetTodoRequest],
) (*connect.Response[todoV1.GetTodoResponse], error) {
	identifier := uuid.FromStringOrNil(request.Msg.GetId())
	if identifier.IsNil() {
		log.Warn().
			Str("provided_id", request.Msg.GetId()).
			Msg("invalid todo ID provided")

		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			ErrInvalidTodoID,
		)
	}

	log.Debug().
		Str("todo_id", identifier.String()).
		Msg("getting todo")

	found, err := t.repo.GetByID(ctx, identifier)
	if err != nil {
		log.Error().
			Err(err).
			Str("todo_id", identifier.String()).
			Msg("failed to get todo")

		return nil, dbErrorToAPI(err, "failed to get todo")
	}

	log.Debug().
		Str("todo_id", identifier.String()).
		Msg("todo retrieved successfully")

	return connect.NewResponse(&todoV1.GetTodoResponse{
		Todo: dbTodoToAPI(found),
	}), nil
}
