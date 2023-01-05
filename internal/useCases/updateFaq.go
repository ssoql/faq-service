package useCases

import (
	"context"
	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/useCases/repositories"
	"github.com/ssoql/faq-service/utils/api_errors"
)

type UpdateFaqUseCase interface {
}

type updateFaqUseCase struct {
	db repositories.DbWrite
}

func NewUpdateFaqUseCase(writeRepository repositories.DbWrite) *updateFaqUseCase {
	return &updateFaqUseCase{db: writeRepository}
}

func (u *updateFaqUseCase) Handle(ctx context.Context, faqID int64, question, answer string) (*entities.Faq, api_errors.ApiError) {
	//return u.db.Update(ctx, &entities.Faq{Id: faqID, Question: question, Answer: answer})
	return nil, nil
}
