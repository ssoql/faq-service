package infrastructure

import (
	"errors"
	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/infrastructure/db"
	"github.com/ssoql/faq-service/utils/api_errors"
	"strings"
)

type faqReadRepository struct {
	db *db.ClientDB
}

func NewFaqReadRepository(db *db.ClientDB) *faqReadRepository {
	return &faqReadRepository{db: db}
}

func (r *faqReadRepository) GetByID(id int64) (*entities.Faq, api_errors.ApiError) {
	faq := &entities.Faq{}

	if err := r.db.Where("id = ?", id).First(faq).Error; err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "record not found") {
			return nil, api_errors.NewNotFoundError("faq with given id does not exists")
		}
		return nil, api_errors.NewInternalServerError(
			"error when tying to fetch faq",
			errors.New("database error"),
		)
	}

	return faq, nil
}
