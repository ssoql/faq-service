package useCases

import (
	"context"
	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/useCases/repositories"
	"github.com/ssoql/faq-service/utils/api_errors"
)

type GetFaqUseCase interface {
	Handle(ctx context.Context, faqID int64) (*entities.Faq, api_errors.ApiError)
}

type getFaqUseCase struct {
	db repositories.DbRead
}

func NewGetFaqUseCase(readRepository repositories.DbRead) *getFaqUseCase {
	return &getFaqUseCase{db: readRepository}
}

func (u *getFaqUseCase) Handle(ctx context.Context, faqID int64) (*entities.Faq, api_errors.ApiError) {
	return u.db.GetByID(faqID)
}
