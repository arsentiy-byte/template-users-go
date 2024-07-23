package user

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"users/internal/model"
	base "users/internal/repository"
	"users/pkg/database"
)

type repository struct {
	db database.Database
}

func NewRepository(db database.Database) base.UserRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) List(ctx context.Context, paging model.Paging) ([]model.User, int64, error) {
	var users []model.User

	query := r.db.GetInstance().Model(&model.User{})

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		WithContext(ctx).
		Order("id DESC").
		Limit(paging.Limit).
		Offset((paging.Offset - 1) * paging.Limit).
		Find(&users).
		Error; err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

func (r *repository) GetById(ctx context.Context, id uint64) (*model.User, error) {
	var user model.User

	if err := r.db.GetInstance().WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r *repository) Create(ctx context.Context, user *model.User) (uint64, error) {
	if err := r.db.GetInstance().WithContext(ctx).Create(user).Error; err != nil {
		return 0, err
	}

	return user.Id, nil
}

func (r *repository) Update(ctx context.Context, user *model.User) error {
	if err := r.db.GetInstance().WithContext(ctx).Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) ForceDelete(ctx context.Context, id uint64) error {
	if err := r.db.GetInstance().WithContext(ctx).Delete(&model.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
