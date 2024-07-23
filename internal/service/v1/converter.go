package v1

import "users/internal/model"

func ConvertPagingDtoToPaging(dto PagingDto) *model.Paging {
	return &model.Paging{
		Limit:  dto.Limit,
		Offset: dto.Offset,
	}
}
