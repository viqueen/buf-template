package api_todo_v1

import "errors"

var (
	ErrInvalidTodoID = errors.New("invalid todo ID")
)

const (
	PageLimitDefault = int32(10)
)
