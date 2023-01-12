package useCases

import (
	"context"
	"github.com/ssoql/faq-service/internal/app/entities"

	"github.com/ssoql/faq-service/internal/useCases/repositories"
	"github.com/ssoql/faq-service/utils/apiErrors"
)

type DeleteFaqUseCase interface {
	Handle(ctx context.Context, faqID int64) apiErrors.ApiError
}

type deleteFaqUseCase struct {
	db repositories.FaqWriteRepository
}

func NewDeleteFaqUseCase(writeRepository repositories.FaqWriteRepository) *deleteFaqUseCase {
	return &deleteFaqUseCase{db: writeRepository}
}

func (u *deleteFaqUseCase) Handle(ctx context.Context, faqID int64) apiErrors.ApiError {
	return u.db.Delete(ctx, &entities.Faq{Id: faqID})
}
