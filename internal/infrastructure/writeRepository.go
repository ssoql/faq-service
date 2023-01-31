package infrastructure

import (
	"context"
	"errors"
	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/infrastructure/db"
	"github.com/ssoql/faq-service/utils/apiErrors"
	"github.com/ssoql/faq-service/utils/cryptoUtils"
	"log"
	"strings"
)

type faqWriteRepository struct {
	db *db.ClientDB
}

func NewFaqWriteRepository(db *db.ClientDB) *faqWriteRepository {
	return &faqWriteRepository{db: db}
}

func (r *faqWriteRepository) Insert(ctx context.Context, faq *entities.Faq) apiErrors.ApiError {
	faq.UniqHash = cryptoUtils.GetMd5(strings.ToLower(faq.Question))

	if err := r.db.WithContext(ctx).Create(faq).Error; err != nil {
		log.Println("error when trying to prepare save faq statement", err.Error())

		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return apiErrors.NewBadRequestError("this question already exists")
		}
		return apiErrors.NewInternalServerError("error when tying to save faq", errors.New("database error"))
	}

	return nil
}

func (r *faqWriteRepository) Update(ctx context.Context, faq *entities.Faq) apiErrors.ApiError {
	faq.UniqHash = cryptoUtils.GetMd5(strings.ToLower(faq.Question))

	if err := r.db.WithContext(ctx).Updates(&faq).Error; err != nil {
		log.Println("error when trying to prepare save faq statement", err.Error())

		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return apiErrors.NewBadRequestError("this question already exists")
		}
		return apiErrors.NewInternalServerError("error when tying to update faq", errors.New("database error"))
	}

	return nil
}

func (r *faqWriteRepository) Delete(ctx context.Context, faq *entities.Faq) apiErrors.ApiError {
	// perform soft delete
	if err := r.db.WithContext(ctx).Delete(faq).Error; err != nil {
		return apiErrors.NewInternalServerError("error when tying to delete faq", errors.New("database error"))
	}

	return nil
}
