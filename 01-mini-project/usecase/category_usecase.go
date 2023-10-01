package usecase

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"
)

type CategoryUseCase interface {
	Create(payload web.CategoryCreateRequest) (web.CategoryResponse, error)
}
