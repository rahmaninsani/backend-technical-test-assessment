package repository

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"gorm.io/gorm"
)

type PostTagRepositoryImpl struct {
	DB *gorm.DB
}

func NewPostTagRepository(db *gorm.DB) PostTagRepository {
	return &PostTagRepositoryImpl{DB: db}
}

func (repository PostTagRepositoryImpl) Save(postTag domain.PostTag) (domain.PostTag, error) {
	if err := repository.DB.Debug().Create(&postTag).Error; err != nil {
		return domain.PostTag{}, err
	}
	return postTag, nil
}

func (repository PostTagRepositoryImpl) Delete(postTag domain.PostTag) error {
	if err := repository.DB.Debug().Delete(&domain.PostTag{}, postTag.PostId).Error; err != nil {
		return err
	}
	
	return nil
}
