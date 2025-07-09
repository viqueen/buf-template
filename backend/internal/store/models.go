package store

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Description string    `gorm:"type:text;not null" json:"description"`
}

func (Todo) TableName() string {
	return "todo"
}

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) CreateTodo(ctx context.Context, todo *Todo) error {
	return r.db.WithContext(ctx).Create(todo).Error
}

func (r *TodoRepository) GetTodo(ctx context.Context, id uuid.UUID) (*Todo, error) {
	var todo Todo
	err := r.db.WithContext(ctx).First(&todo, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) ListTodos(ctx context.Context, limit, offset int32) ([]*Todo, error) {
	var todos []*Todo
	err := r.db.WithContext(ctx).Limit(int(limit)).Offset(int(offset)).Find(&todos).Error
	return todos, err
}