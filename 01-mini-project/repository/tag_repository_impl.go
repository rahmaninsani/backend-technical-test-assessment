package repository

import (
	"errors"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"gorm.io/gorm"
)

type TagRepositoryImpl struct {
	DB *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &TagRepositoryImpl{DB: db}
}

func (repository TagRepositoryImpl) Save(tag domain.Tag) (domain.Tag, error) {
	if err := repository.DB.Debug().Create(&tag).Error; err != nil {
		return domain.Tag{}, err
	}
	return tag, nil
}

func (repository TagRepositoryImpl) FindOne(tag domain.Tag) (domain.Tag, error) {
	if err := repository.DB.Debug().
		Where(&tag).
		First(&tag).
		Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Tag{}, err
	}
	
	return tag, nil
}
