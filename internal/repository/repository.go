package repository

import (
	"context"
	"users/internal/model"
)

type UserRepository interface {
	List(ctx context.Context, paging model.Paging) ([]model.User, int64, error)
	GetById(ctx context.Context, id uint) (*model.User, error)
	Create(ctx context.Context, user *model.User) (uint, error)
	Update(ctx context.Context, user *model.User) error
	ForceDelete(ctx context.Context, id uint) error
}
