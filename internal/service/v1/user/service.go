package user

import (
	"context"
	"users/internal/repository"
	v1 "users/internal/service/v1"
)

type service struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) v1.UserService {
	return &service{repo: repo}
}

func (s *service) List(ctx context.Context, dto v1.PagingDto) ([]Dto, int64, error) {
	users, count, err := s.repo.List(ctx, *v1.ConvertPagingDtoToPaging(dto))
	if err != nil {
		return nil, 0, err
	}

	var dtos []Dto
	for _, user := range users {
		dtos = append(dtos, *convertUserToDto(user))
	}

	return dtos, count, nil
}

func (s *service) GetById(ctx context.Context, id uint64) (*Dto, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertUserToDto(*user), nil
}

func (s *service) Create(ctx context.Context, dto CreateDto) (uint64, error) {
	user := convertCreateDtoToUser(dto)

	return s.repo.Create(ctx, user)
}

func (s *service) Update(ctx context.Context, dto UpdateDto) error {
	user := convertUpdateDtoToUser(dto)

	return s.repo.Update(ctx, user)
}

func (s *service) SoftDelete(ctx context.Context, id uint64) error {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	user.SoftDelete()

	return s.repo.Update(ctx, user)
}
