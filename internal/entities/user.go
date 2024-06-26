package entities

import "time"

type User struct {
	Id        uint       `gorm:"primaryKey"`
	FirstName string     `gorm:"type:varchar(255);column:first_name;not null"`
	LastName  string     `gorm:"type:varchar(255);column:last_name;not null"`
	Email     string     `gorm:"type:varchar(255);column:email;unique;not null"`
	Password  string     `gorm:"type:varchar(255);column:password;not null"`
	CreatedAt time.Time  `gorm:"type:timestamptz;column:created_at;not null"`
	UpdatedAt time.Time  `gorm:"type:timestamptz;column:created_at;not null"`
	DeletedAt *time.Time `gorm:"type:timestamptz;column:deleted_at"`
}
