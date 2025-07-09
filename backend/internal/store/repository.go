package store

import (
	"context"

	"github.com/gofrs/uuid"
)

type ListOptions struct {
	Limit  int32
	Offset int32
}

type Repository[T any] interface {
	Create(ctx context.Context, entity *T) error
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*T, error)
	List(ctx context.Context, opts ListOptions) ([]*T, error)
}
