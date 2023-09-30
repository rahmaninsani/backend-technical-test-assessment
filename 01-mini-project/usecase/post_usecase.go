package usecase

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"
)

type PostUseCase interface {
	Create(payload web.PostCreateRequest, user domain.User) (web.PostResponse, error)
}
