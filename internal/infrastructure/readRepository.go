package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"

	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/infrastructure/db"
	"github.com/ssoql/faq-service/utils/apiErrors"
)

type faqReadRepository struct {
	db *db.ClientDB
}

func NewFaqReadRepository(db *db.ClientDB) *faqReadRepository {
	return &faqReadRepository{db: db}
}

func (r *faqReadRepository) GetByID(ctx context.Context, id int64) (*entities.Faq, apiErrors.ApiError) {
	faq := &entities.Faq{}

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(faq).Error; err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "record not found") {
			return nil, apiErrors.NewNotFoundError("faq with given id does not exists")
		}
		return nil, apiErrors.NewInternalServerError(
			"error when tying to fetch faq",
			fmt.Errorf("database error: %s", err.Error()),
		)
	}

	return faq, nil
}

func (r *faqReadRepository) GetAll(ctx context.Context, page, pageSize int) (*entities.Faqs, apiErrors.ApiError) {
	var faqs = &entities.Faqs{}

	if err := r.db.WithContext(ctx).Scopes(Paginate(page, pageSize)).Find(faqs).Error; err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "record not found") {
			return faqs, apiErrors.NewNotFoundError("there is no faqs in the DB")
		}
		return faqs, apiErrors.NewInternalServerError("error when tying to fetch faqs", err)
	}
	return faqs, nil
}

func (r *faqReadRepository) Exists(ctx context.Context, id int64) (bool, apiErrors.ApiError) {
	var exists bool
	err := r.db.WithContext(ctx).Model(&entities.Faq{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).
		Error

	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "record not found") {
			return false, apiErrors.NewNotFoundError("faq with given id does not exists")
		}
		return false, apiErrors.NewInternalServerError(
			"error when tying to fetch faq",
			errors.New("database error"),
		)
	}

	return exists, nil
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}
