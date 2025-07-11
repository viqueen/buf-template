package api_todo_v1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
	todoV1 "github.com/viqueen/buf-template/api/go-sdk/todo/v1"
	"github.com/viqueen/buf-template/backend/internal/store"
)

func (t todoService) CreateTodo(
	ctx context.Context,
	request *connect.Request[todoV1.CreateTodoRequest],
) (*connect.Response[todoV1.CreateTodoResponse], error) {
	id := uuid.Must(uuid.NewV4())

	log.Debug().
		Str("todo_id", id.String()).
		Str("description", request.Msg.GetDescription()).
		Msg("creating todo")

	todo := &store.Todo{
		ID:          id,
		Description: request.Msg.GetDescription(),
	}

	err := t.repo.Create(ctx, todo)
	if err != nil {
		log.Error().
			Err(err).
			Str("todo_id", id.String()).
			Msg("failed to create todo")
		return nil, dbErrorToAPI(err, "failed to create todo")
	}

	log.Info().
		Str("todo_id", id.String()).
		Msg("todo created successfully")

	return connect.NewResponse(&todoV1.CreateTodoResponse{
		Todo: dbTodoToAPI(todo),
	}), nil
}
