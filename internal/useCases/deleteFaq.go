package useCases

import (
	"context"
	"github.com/ssoql/faq-service/internal/app/entities"

	"github.com/ssoql/faq-service/internal/useCases/repositories"
	"github.com/ssoql/faq-service/utils/api_errors"
)

type DeleteFaqUseCase interface {
	Handle(ctx context.Context, faqID int64) api_errors.ApiError
}

type deleteFaqUseCase struct {
	db repositories.DbWrite
}

func NewDeleteFaqUseCase(writeRepository repositories.DbWrite) *deleteFaqUseCase {
	return &deleteFaqUseCase{db: writeRepository}
}

func (u *deleteFaqUseCase) Handle(ctx context.Context, faqID int64) api_errors.ApiError {
	return u.db.Delete(ctx, &entities.Faq{Id: faqID})
}
