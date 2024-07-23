package user

import "users/internal/model"

func convertUserToDto(user model.User) *Dto {
	return &Dto{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func convertDtoToUser(dto Dto) *model.User {
	return &model.User{
		Id:        dto.Id,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		DeletedAt: dto.DeletedAt,
	}
}

func convertCreateDtoToUser(dto CreateDto) *model.User {
	return &model.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  dto.Password,
	}
}

func convertUpdateDtoToUser(dto UpdateDto) *model.User {
	return &model.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
	}
}
