package useCases

import (
	"context"
	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/useCases/repositories"
	"github.com/ssoql/faq-service/utils/apiErrors"
)

type UpdateFaqUseCase interface {
	Handle(ctx context.Context, faqID int64, question, answer string) (*entities.Faq, apiErrors.ApiError)
}

type updateFaqUseCase struct {
	dbWrite repositories.FaqWriteRepository
	dbRead  repositories.FaqReadRepository
}

func NewUpdateFaqUseCase(writeRepository repositories.FaqWriteRepository, readRepository repositories.FaqReadRepository) *updateFaqUseCase {
	return &updateFaqUseCase{
		dbWrite: writeRepository,
		dbRead:  readRepository,
	}
}

func (u *updateFaqUseCase) Handle(ctx context.Context, faqID int64, question, answer string) (*entities.Faq, apiErrors.ApiError) {
	shutdownCtx := handleShutdown(ctx)

	exist, err := u.dbRead.Exists(shutdownCtx, faqID)
	if err != nil {
		return &entities.Faq{}, err
	}

	if !exist {
		return &entities.Faq{}, apiErrors.NewNotFoundError("faq with given id does not exist")
	}

	faq := &entities.Faq{Id: faqID, Question: question, Answer: answer}

	if err := u.dbWrite.Update(shutdownCtx, faq); err != nil {
		return nil, err
	}

	return faq, nil
}
