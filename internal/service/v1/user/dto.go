package user

import "time"

type CreateDto struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UpdateDto struct {
	FirstName string
	LastName  string
	Email     string
}

type Dto struct {
	Id        uint64
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
