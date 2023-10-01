package repository

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
)

type PostTagRepository interface {
	Save(postTag domain.PostTag) (domain.PostTag, error)
	Delete(postTag domain.PostTag) error
}
