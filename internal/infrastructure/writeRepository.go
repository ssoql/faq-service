package infrastructure

import (
	"context"
	"errors"
	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/infrastructure/db"
	"github.com/ssoql/faq-service/utils/api_errors"
	"github.com/ssoql/faq-service/utils/crypto_utils"
	"log"
	"strings"
)

type faqWriteRepository struct {
	db *db.ClientDB
}

func NewFaqWriteRepository(db *db.ClientDB) *faqWriteRepository {
	return &faqWriteRepository{db: db}
}

func (r *faqWriteRepository) Insert(ctx context.Context, faq *entities.Faq) api_errors.ApiError {
	faq.UniqHash = crypto_utils.GetMd5(strings.ToLower(faq.Question))
	if err := r.db.Create(faq).Error; err != nil {
		log.Println("error when trying to prepare save faq statement", err.Error())

		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return api_errors.NewBadRequestError("this question already exists")
		}
		return api_errors.NewInternalServerError("error when tying to save faq", errors.New("database error"))
	}

	return nil
}

func (r *faqWriteRepository) Update(ctx context.Context, faq *entities.Faq) api_errors.ApiError {
	if err := r.db.Updates(&faq).Error; err != nil {
		log.Println("error when trying to prepare save faq statement", err.Error())

		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return api_errors.NewBadRequestError("this question already exists")
		}
		return api_errors.NewInternalServerError("error when tying to save faq", errors.New("database error"))
	}

	return nil
}

func (r *faqWriteRepository) Delete(ctx context.Context, faq *entities.Faq) api_errors.ApiError {
	// perform soft delete
	if err := r.db.Delete(faq).Error; err != nil {
		return api_errors.NewInternalServerError("error when tying to delete faq", errors.New("database error"))
	}

	return nil
}
