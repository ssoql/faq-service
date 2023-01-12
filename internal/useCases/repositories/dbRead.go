//go:generate mockery --output mocks --name FaqReadRepository --exported --filename faqReadRepository.go
package repositories

import (
	"context"

	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/utils/apiErrors"
)

type FaqReadRepository interface {
	GetByID(ctx context.Context, id int64) (*entities.Faq, apiErrors.ApiError)
	GetAll(ctx context.Context, page, pageSize int) (*entities.Faqs, apiErrors.ApiError)
	Exists(ctx context.Context, id int64) (bool, apiErrors.ApiError)
}
