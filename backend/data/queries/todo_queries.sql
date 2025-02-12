-- name: CreateTodo :one
-- CreateTodo creates a new todo.
INSERT INTO todo (id,  description)
VALUES (@id,  @description)
RETURNING *;

-- name: GetTodo :one
-- GetTodo retrieves a todo by id.
SELECT * FROM todo WHERE id = @id;

-- name: ListTodos :many
-- ListTodos lists all todo.
SELECT * FROM todo
LIMIT @page_limit OFFSET @page_offset;