package repository

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
)

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
}
