package repository

import (
	"context"
	"users/internal/model"
)

type UserRepository interface {
	List(ctx context.Context, paging model.Paging) ([]model.User, int64, error)
	GetById(ctx context.Context, id uint64) (*model.User, error)
	Create(ctx context.Context, user *model.User) (uint64, error)
	Update(ctx context.Context, user *model.User) error
	ForceDelete(ctx context.Context, id uint64) error
}
