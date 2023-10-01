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

func (repository CategoryRepositoryImpl) Save(category domain.Category) (domain.Category, error) {
	if err := repository.DB.Debug().Create(&category).Error; err != nil {
		return domain.Category{}, err
	}
	return category, nil
}

func (repository CategoryRepositoryImpl) FindAll() ([]domain.Category, error) {
	var categories []domain.Category
	
	if err := repository.DB.Debug().
		Find(&categories).
		Error; err != nil {
		return categories, err
	}
	
	return categories, nil
}

func (repository CategoryRepositoryImpl) FindOne(category domain.Category) (domain.Category, error) {
	if err := repository.DB.Debug().
		Where(&category).
		First(&category).
		Error; err != nil {
		return domain.Category{}, err
	}
	
	return category, nil
}
