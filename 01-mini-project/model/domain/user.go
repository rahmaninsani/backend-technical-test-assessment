package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Username  string    `gorm:"type:varchar(100);not null;unique"`
	Email     string    `gorm:"type:varchar(100);not null;unique"`
	Password  string    `gorm:"type:varchar(60);not null"`
	Avatar    string    `gorm:"type:varchar(45);not null;default:'avatar_default.png'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	
	Posts []Post
}
