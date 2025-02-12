package api_todo_v1

import (
	todoV1 "api/go-sdk/todo/v1"
	"backend/internal/store/gendb"
	"connectrpc.com/connect"
	"database/sql"
	"errors"
)

func dbErrorToApi(err error, msg string) *connect.Error {
	if errors.Is(err, sql.ErrNoRows) {
		return connect.NewError(connect.CodeNotFound, err)
	}
	wrapped := errors.Join(errors.New(msg), err)
	return connect.NewError(connect.CodeInternal, wrapped)
}

func dbTodoToApi(todo *gendb.Todo) *todoV1.Todo {
	return &todoV1.Todo{
		Id:          todo.ID.String(),
		Description: todo.Description,
	}
}

func nonZeroOrDefaultInt32(i int32, def int32) int32 {
	if i == 0 {
		return def
	}
	return i
}
