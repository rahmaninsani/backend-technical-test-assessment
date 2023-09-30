package domain

import (
	"github.com/google/uuid"
	"time"
)

type PostTag struct {
	PostId    uuid.UUID `gorm:"type:uuid;primaryKey;onUpdate:CASCADE;onDelete:CASCADE"`
	TagId     uuid.UUID `gorm:"type:uuid;primaryKey;onUpdate:CASCADE;onDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
