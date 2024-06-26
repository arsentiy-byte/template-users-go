package model

type Paging struct {
	Offset int
	Limit  int
}

func NewPaging(offset int, limit int) *Paging {
	return &Paging{
		Offset: offset,
		Limit:  limit,
	}
}
