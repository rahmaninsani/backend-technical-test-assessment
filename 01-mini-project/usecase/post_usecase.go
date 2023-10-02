package usecase

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"
)

type PostUseCase interface {
	Create(payload web.PostCreateRequest, user domain.User) (web.PostResponse, error)
	Update(payload web.PostUpdateRequest, user domain.User) (web.PostResponse, error)
	Delete(payload web.PostDeleteRequest, user domain.User) error
	FindOne(payload web.PostFindOneRequest) (web.PostResponse, error)
}
