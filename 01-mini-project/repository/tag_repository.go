package repository

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
)

type TagRepository interface {
	Save(tag domain.Tag) (domain.Tag, error)
	FindOne(tag domain.Tag) (domain.Tag, error)
}
