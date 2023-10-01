package domain

import (
	"github.com/google/uuid"
	"time"
)

type Tag struct {
	Id        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	
	PostTags []PostTag
}
