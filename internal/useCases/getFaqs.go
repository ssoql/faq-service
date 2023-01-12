package useCases

import (
	"context"

	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/useCases/repositories"
	"github.com/ssoql/faq-service/utils/apiErrors"
)

type GetFaqsUseCase interface {
	Handle(ctx context.Context, pageNumber, pageSize int) (*entities.Faqs, apiErrors.ApiError)
}

type getFaqsUseCase struct {
	db repositories.FaqReadRepository
}

func NewGetFaqsUseCase(readRepository repositories.FaqReadRepository) *getFaqsUseCase {
	return &getFaqsUseCase{db: readRepository}
}

func (u *getFaqsUseCase) Handle(ctx context.Context, pageNumber, pageSize int) (*entities.Faqs, apiErrors.ApiError) {
	return u.db.GetAll(ctx, pageNumber, pageSize)
}
