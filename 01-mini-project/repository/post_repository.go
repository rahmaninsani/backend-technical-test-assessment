package repository

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
)

type PostRepository interface {
	Save(post domain.Post) (domain.Post, error)
	FindOne(post domain.Post) (domain.Post, error)
	Update(post domain.Post) (domain.Post, error)
	Delete(post domain.Post) error
}
