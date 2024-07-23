package v1

import (
	"context"
	"users/internal/service/v1/user"
)

type UserService interface {
	List(ctx context.Context, paging PagingDto) ([]user.Dto, int64, error)
	GetById(ctx context.Context, id uint64) (*user.Dto, error)
	Create(ctx context.Context, dto user.CreateDto) (uint64, error)
	Update(ctx context.Context, dto user.UpdateDto) error
	SoftDelete(ctx context.Context, id uint64) error
}
