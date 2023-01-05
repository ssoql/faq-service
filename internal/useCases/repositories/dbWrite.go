package repositories

import (
	"context"

	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/utils/api_errors"
)

type DbWrite interface {
	Insert(ctx context.Context, faq *entities.Faq) api_errors.ApiError
	Update(ctx context.Context, faq *entities.Faq) api_errors.ApiError
	Delete(ctx context.Context, faq *entities.Faq) api_errors.ApiError
}
