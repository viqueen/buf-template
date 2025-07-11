package api_todo_v1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/rs/zerolog/log"
	todoV1 "github.com/viqueen/buf-template/api/go-sdk/todo/v1"
	"github.com/viqueen/buf-template/backend/internal/store"
)

func (t todoService) ListTodos(
	ctx context.Context,
	request *connect.Request[todoV1.ListTodosRequest],
) (*connect.Response[todoV1.ListTodosResponse], error) {
	limit := nonZeroOrDefaultInt32(request.Msg.GetPageLimit(), PageLimitDefault)
	offset := request.Msg.GetPageOffset()

	log.Debug().
		Int32("limit", limit).
		Int32("offset", offset).
		Msg("listing todos")

	todos, err := t.repo.List(ctx, store.ListOptions{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		log.Error().
			Err(err).
			Int32("limit", limit).
			Int32("offset", offset).
			Msg("failed to list todos")

		return nil, dbErrorToAPI(err, "failed to list todos")
	}

	apiTodos := make([]*todoV1.Todo, 0, len(todos))
	for index, todo := range todos {
		apiTodos[index] = dbTodoToAPI(todo)
	}

	log.Info().
		Int("count", len(todos)).
		Int32("limit", limit).
		Int32("offset", offset).
		Msg("todos listed successfully")

	return connect.NewResponse(&todoV1.ListTodosResponse{
		Todos: apiTodos,
	}), nil
}
