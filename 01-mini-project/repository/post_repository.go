package repository

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
)

type PostRepository interface {
	Save(user domain.Post) (domain.Post, error)
}
