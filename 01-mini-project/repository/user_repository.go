package repository

import (
	"github.com/google/uuid"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
)

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
	FindOneByEmail(email string) (domain.User, error)
	FindOneByUserId(userId uuid.UUID) (domain.User, error)
}
