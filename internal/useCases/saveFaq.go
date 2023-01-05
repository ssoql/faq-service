package useCases

import (
	"context"
	"github.com/ssoql/faq-service/internal/app/entities"

	"github.com/ssoql/faq-service/internal/useCases/repositories"
	"github.com/ssoql/faq-service/utils/api_errors"
)

type SaveFaqUseCase interface {
	Handle(ctx context.Context, question, answer string) (*entities.Faq, api_errors.ApiError)
}

type saveFaqUseCase struct {
	db repositories.DbWrite
}

func NewCreateFaqUseCase(writeRepository repositories.DbWrite) *saveFaqUseCase {
	return &saveFaqUseCase{db: writeRepository}
}

func (u *saveFaqUseCase) Handle(ctx context.Context, question, answer string) (*entities.Faq, api_errors.ApiError) {
	newFaq := &entities.Faq{
		Question: question,
		Answer:   answer,
	}

	if err := u.db.Insert(ctx, newFaq); err != nil {
		return nil, err
	}

	return newFaq, nil
}
