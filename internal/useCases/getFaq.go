package useCases

import (
	"context"

	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/useCases/repositories"
	"github.com/ssoql/faq-service/utils/apiErrors"
)

type GetFaqUseCase interface {
	Handle(ctx context.Context, faqID int64) (*entities.Faq, apiErrors.ApiError)
}

type getFaqUseCase struct {
	db repositories.FaqReadRepository
}

func NewGetFaqUseCase(readRepository repositories.FaqReadRepository) *getFaqUseCase {
	return &getFaqUseCase{db: readRepository}
}

func (u *getFaqUseCase) Handle(ctx context.Context, faqID int64) (*entities.Faq, apiErrors.ApiError) {
	return u.db.GetByID(ctx, faqID)
}
