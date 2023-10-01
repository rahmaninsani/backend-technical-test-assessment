package domain

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	Id        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	
	Posts []Post
}
