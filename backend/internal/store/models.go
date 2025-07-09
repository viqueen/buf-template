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

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) Repository[Todo] {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(ctx context.Context, todo *Todo) error {
	return r.db.WithContext(ctx).Create(todo).Error
}

func (r *todoRepository) Update(ctx context.Context, todo *Todo) error {
	return r.db.WithContext(ctx).Save(todo).Error
}

func (r *todoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&Todo{}, "id = ?", id).Error
}

func (r *todoRepository) GetByID(ctx context.Context, id uuid.UUID) (*Todo, error) {
	var todo Todo
	err := r.db.WithContext(ctx).First(&todo, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) List(ctx context.Context, opts ListOptions) ([]*Todo, error) {
	var todos []*Todo
	err := r.db.WithContext(ctx).Limit(int(opts.Limit)).Offset(int(opts.Offset)).Find(&todos).Error
	return todos, err
}