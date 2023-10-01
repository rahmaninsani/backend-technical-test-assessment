package repository

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
)

type CategoryRepository interface {
	Save(user domain.Category) (domain.Category, error)
}
