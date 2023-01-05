package repositories

import (
	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/utils/api_errors"
)

type DbRead interface {
	GetByID(id int64) (*entities.Faq, api_errors.ApiError)
}
