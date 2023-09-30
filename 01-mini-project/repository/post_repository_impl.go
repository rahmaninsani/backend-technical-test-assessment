package repository

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{DB: db}
}

func (repository PostRepositoryImpl) Save(post domain.Post) (domain.Post, error) {
	if err := repository.DB.Debug().Create(&post).Error; err != nil {
		return domain.Post{}, err
	}
	return post, nil
}
