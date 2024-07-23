package model

import (
	"errors"
	"time"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	Id        uint64     `gorm:"primaryKey"`
	FirstName string     `gorm:"type:varchar(255);column:first_name;not null"`
	LastName  string     `gorm:"type:varchar(255);column:last_name;not null"`
	Email     string     `gorm:"type:varchar(255);column:email;unique;not null"`
	Password  string     `gorm:"type:varchar(255);column:password;not null"`
	CreatedAt time.Time  `gorm:"type:timestamptz;column:created_at;not null"`
	UpdatedAt time.Time  `gorm:"type:timestamptz;column:created_at;not null"`
	DeletedAt *time.Time `gorm:"type:timestamptz;column:deleted_at"`
}

func NewUser(
	firstName string,
	lastName string,
	email string,
	password string,
) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

func (u *User) Update(
	firstName string,
	lastName string,
	email string,
) {
	u.FirstName = firstName
	u.LastName = lastName
	u.Email = email
	u.UpdatedAt = time.Now().UTC()
}

func (u *User) SoftDelete() {
	now := time.Now().UTC()

	u.DeletedAt = &now
}
