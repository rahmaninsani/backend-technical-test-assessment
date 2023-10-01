package repository

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
)

type CategoryRepository interface {
	Save(category domain.Category) (domain.Category, error)
	FindAll() ([]domain.Category, error)
	FindOne(category domain.Category) (domain.Category, error)
}
