//go:generate mockery --output mocks --name FaqWriteRepository --exported --filename faqWriteRepository.go
package repositories

import (
	"context"

	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/utils/apiErrors"
)

type FaqWriteRepository interface {
	Insert(ctx context.Context, faq *entities.Faq) apiErrors.ApiError
	Update(ctx context.Context, faq *entities.Faq) apiErrors.ApiError
	Delete(ctx context.Context, faq *entities.Faq) apiErrors.ApiError
}
