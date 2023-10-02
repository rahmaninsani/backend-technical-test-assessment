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

func (repository PostRepositoryImpl) FindOne(post domain.Post) (domain.Post, error) {
	if err := repository.DB.Debug().
		Where(&post).
		First(&post).
		Error; err != nil {
		return domain.Post{}, err
	}
	
	return post, nil
}

func (repository PostRepositoryImpl) Update(post domain.Post) (domain.Post, error) {
	if err := repository.DB.Debug().Save(&post).Error; err != nil {
		return domain.Post{}, err
	}
	return post, nil
}

func (repository PostRepositoryImpl) Delete(post domain.Post) error {
	if err := repository.DB.Debug().Delete(&domain.Post{}, post.Id).Error; err != nil {
		return err
	}
	
	return nil
}
