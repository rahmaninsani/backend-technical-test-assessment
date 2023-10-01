package repository

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
)

type PostTagRepository interface {
	Save(category domain.PostTag) (domain.PostTag, error)
}
