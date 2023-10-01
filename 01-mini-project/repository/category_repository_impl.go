package repository

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}

func (repository CategoryRepositoryImpl) Save(post domain.Category) (domain.Category, error) {
	if err := repository.DB.Debug().Create(&post).Error; err != nil {
		return domain.Category{}, err
	}
	return post, nil
}
