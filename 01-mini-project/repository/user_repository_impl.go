package repository

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (repository UserRepositoryImpl) Save(user domain.User) (domain.User, error) {
	if err := repository.DB.Debug().Create(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repository UserRepositoryImpl) FindOne(user domain.User) (domain.User, error) {
	if err := repository.DB.Debug().
		Where(&user).
		First(&user).
		Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}
