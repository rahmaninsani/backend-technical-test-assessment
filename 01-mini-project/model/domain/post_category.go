package domain

import (
	"github.com/google/uuid"
	"time"
)

type PostCategory struct {
	PostId     uuid.UUID `gorm:"type:uuid;primaryKey;onUpdate:CASCADE;onDelete:CASCADE"`
	CategoryId uuid.UUID `gorm:"type:uuid;primaryKey;onUpdate:CASCADE;onDelete:CASCADE"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
