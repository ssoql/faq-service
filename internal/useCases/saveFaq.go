package useCases

import (
	"context"
	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/global"
	"os"

	"github.com/ssoql/faq-service/internal/useCases/repositories"
	"github.com/ssoql/faq-service/utils/apiErrors"
)

type SaveFaqUseCase interface {
	Handle(ctx context.Context, question, answer string) (*entities.Faq, apiErrors.ApiError)
}

type saveFaqUseCase struct {
	db repositories.FaqWriteRepository
}

func NewCreateFaqUseCase(writeRepository repositories.FaqWriteRepository) *saveFaqUseCase {
	return &saveFaqUseCase{db: writeRepository}
}

func (u *saveFaqUseCase) Handle(ctx context.Context, question, answer string) (*entities.Faq, apiErrors.ApiError) {
	newFaq := &entities.Faq{
		Question: question,
		Answer:   answer,
	}

	shutdownCtx := u.handleShutdown(ctx)

	if err := u.db.Insert(shutdownCtx, newFaq); err != nil {
		return nil, err
	}

	return newFaq, nil
}

// handleShutdown returns context that will be cancelled if shutdown will occur
func (u *saveFaqUseCase) handleShutdown(ctx context.Context) context.Context {
	shutdownChan := ctx.Value(global.ShutdownSignal).(chan os.Signal)
	shutdownCtx, cancel := context.WithCancel(ctx)

	go func(cancel context.CancelFunc) {
		<-shutdownChan
		cancel()
	}(cancel)

	return shutdownCtx
}
