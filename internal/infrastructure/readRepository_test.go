package infrastructure

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/infrastructure/db/mocks"
	"github.com/ssoql/faq-service/utils/apiErrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"strconv"
	"testing"
)

var (
	TestFaqID       = int64(1)
	TestFaqHash     = "35179a54ea587953021400eb0cd23201"
	TestFaqQuestion = "how are you?"
	TestFaqAnswer   = "i am fine and you?"
)

func TestFaqReadRepository_GetByID(t *testing.T) {
	type args struct {
		faqID int64
	}
	tests := []struct {
		name         string
		args         args
		mockQueries  func(t *testing.T, mockedDb sqlmock.Sqlmock)
		want         *entities.Faq
		assertResult func(t *testing.T, want, result *entities.Faq, err apiErrors.ApiError)
	}{
		{
			name: "success",
			args: args{
				faqID: TestFaqID,
			},
			mockQueries: func(t *testing.T, mockedDb sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "uniq_hash", "question", "answer"}).
					AddRow(
						TestFaqID,
						TestFaqHash,
						TestFaqQuestion,
						TestFaqAnswer)

				mockedDb.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM ` + "`faqs`" + ` WHERE id = ? ORDER BY ` + "`faqs`.`id`" + ` LIMIT 1`)).
					WithArgs(1).
					WillReturnRows(rows)
			},
			want: &entities.Faq{
				Id:       TestFaqID,
				UniqHash: TestFaqHash,
				Question: TestFaqQuestion,
				Answer:   TestFaqAnswer,
			},
			assertResult: func(t *testing.T, want, result *entities.Faq, err apiErrors.ApiError) {
				require.NoError(t, err)
				assert.Equal(t, want.Id, result.Id)
				assert.NotEmpty(t, result.UniqHash)
				assert.NotEmpty(t, result.Question)
				assert.NotEmpty(t, result.Answer)
			},
		},
		{
			name: "empty-result",
			args: args{
				faqID: TestFaqID,
			},
			mockQueries: func(t *testing.T, mockedDb sqlmock.Sqlmock) {
				mockedDb.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM ` + "`faqs`" + ` WHERE id = ? ORDER BY ` + "`faqs`.`id`" + ` LIMIT 1`)).
					WithArgs(1).
					WillReturnError(errors.New("record not found"))
			},
			want: &entities.Faq{},
			assertResult: func(t *testing.T, want, result *entities.Faq, err apiErrors.ApiError) {
				require.ErrorContains(t, err, "faq with given id does not exists")
			},
		},
		{
			name: "database-failure",
			args: args{
				faqID: TestFaqID,
			},
			mockQueries: func(t *testing.T, mockedDb sqlmock.Sqlmock) {
				mockedDb.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM ` + "`faqs`" + ` WHERE id = ? ORDER BY ` + "`faqs`.`id`" + ` LIMIT 1`)).
					WithArgs(1).
					WillReturnError(errors.New("database error"))
			},
			want: &entities.Faq{},
			assertResult: func(t *testing.T, want, result *entities.Faq, err apiErrors.ApiError) {
				require.ErrorContains(t, err, "error when tying to fetch faq")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockedDB, mockedClient, err := mocks.MockClientDb()

			defer func() {
				dbInstance, err := mockedClient.DB.DB()
				if err != nil {
					_ = dbInstance.Close()
				}
			}()

			require.NoError(t, err)
			tt.mockQueries(t, mockedDB)

			result, repoErr := NewFaqReadRepository(mockedClient).GetByID(context.TODO(), tt.args.faqID)

			tt.assertResult(t, tt.want, result, repoErr)
		})
	}
}

func TestFaqReadRepository_GetAll(t *testing.T) {
	var getAllLimit = 10
	type args struct {
		faqID int64
	}
	tests := []struct {
		name         string
		args         args
		mockQueries  func(t *testing.T, mockedDb sqlmock.Sqlmock)
		want         *entities.Faqs
		assertResult func(t *testing.T, want, result *entities.Faqs, err apiErrors.ApiError)
	}{
		{
			name: "success",
			args: args{
				faqID: TestFaqID,
			},
			mockQueries: func(t *testing.T, mockedDb sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "uniq_hash", "question", "answer"}).
					AddRow(
						TestFaqID,
						TestFaqHash,
						TestFaqQuestion,
						TestFaqAnswer)

				mockedDb.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM ` + "`faqs`" + ` LIMIT ` + strconv.Itoa(getAllLimit))).
					WillReturnRows(rows)
			},
			want: &entities.Faqs{{
				Id:       TestFaqID,
				UniqHash: TestFaqHash,
				Question: TestFaqQuestion,
				Answer:   TestFaqAnswer,
			}},
			assertResult: func(t *testing.T, want, result *entities.Faqs, err apiErrors.ApiError) {
				require.NoError(t, err)
				assert.Equal(t, len(*want), len(*result))
				assert.NotEmpty(t, (*result)[0].UniqHash)
				assert.NotEmpty(t, (*result)[0].Question)
				assert.NotEmpty(t, (*result)[0].Answer)
			},
		},
		{
			name: "empty-result",
			args: args{
				faqID: TestFaqID,
			},
			mockQueries: func(t *testing.T, mockedDb sqlmock.Sqlmock) {
				mockedDb.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM ` + "`faqs`" + ` LIMIT ` + strconv.Itoa(getAllLimit))).
					WillReturnError(errors.New("record not found"))
			},
			want: &entities.Faqs{},
			assertResult: func(t *testing.T, want, result *entities.Faqs, err apiErrors.ApiError) {
				require.ErrorContains(t, err, "there is no faqs in the DB")
			},
		},
		{
			name: "database-failure",
			args: args{
				faqID: TestFaqID,
			},
			mockQueries: func(t *testing.T, mockedDb sqlmock.Sqlmock) {
				mockedDb.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM ` + "`faqs`" + ` LIMIT ` + strconv.Itoa(getAllLimit))).
					WillReturnError(errors.New("database error"))
			},
			want: &entities.Faqs{},
			assertResult: func(t *testing.T, want, result *entities.Faqs, err apiErrors.ApiError) {
				require.ErrorContains(t, err, "error when tying to fetch faq")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockedDB, mockedClient, err := mocks.MockClientDb()

			defer func() {
				dbInstance, err := mockedClient.DB.DB()
				if err != nil {
					_ = dbInstance.Close()
				}
			}()

			require.NoError(t, err)
			tt.mockQueries(t, mockedDB)

			result, repoErr := NewFaqReadRepository(mockedClient).GetAll(context.TODO(), 1, 10)

			tt.assertResult(t, tt.want, result, repoErr)
		})
	}
}

func TestFaqReadRepository_Exist(t *testing.T) {
	type args struct {
		faqID int64
	}
	tests := []struct {
		name         string
		args         args
		mockQueries  func(t *testing.T, mockedDb sqlmock.Sqlmock)
		want         bool
		assertResult func(t *testing.T, want, result bool, err apiErrors.ApiError)
	}{
		{
			name: "success",
			args: args{
				faqID: TestFaqID,
			},
			mockQueries: func(t *testing.T, mockedDb sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"count(*) > 0"}).
					AddRow(true)

				mockedDb.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) > 0  FROM ` + "`faqs`" + ` WHERE id = ?`)).
					WithArgs(1).
					WillReturnRows(rows)
			},
			want: true,
			assertResult: func(t *testing.T, want, result bool, err apiErrors.ApiError) {
				require.NoError(t, err)
				assert.Equal(t, want, result)
			},
		},
		{
			name: "empty-result",
			args: args{
				faqID: TestFaqID,
			},
			mockQueries: func(t *testing.T, mockedDb sqlmock.Sqlmock) {
				mockedDb.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) > 0  FROM ` + "`faqs`" + ` WHERE id = ?`)).
					WithArgs(1).
					WillReturnError(errors.New("record not found"))
			},
			want: false,
			assertResult: func(t *testing.T, want, result bool, err apiErrors.ApiError) {
				assert.ErrorContains(t, err, "faq with given id does not exists")
				assert.Equal(t, want, result)
			},
		},
		{
			name: "database-failure",
			args: args{
				faqID: TestFaqID,
			},
			mockQueries: func(t *testing.T, mockedDb sqlmock.Sqlmock) {
				mockedDb.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) > 0  FROM ` + "`faqs`" + ` WHERE id = ?`)).
					WithArgs(1).
					WillReturnError(errors.New("database error"))
			},
			want: false,
			assertResult: func(t *testing.T, want, result bool, err apiErrors.ApiError) {
				assert.ErrorContains(t, err, "error when tying to fetch faq")
				assert.Equal(t, want, result)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockedDB, mockedClient, err := mocks.MockClientDb()

			defer func() {
				dbInstance, err := mockedClient.DB.DB()
				if err != nil {
					_ = dbInstance.Close()
				}
			}()

			require.NoError(t, err)
			tt.mockQueries(t, mockedDB)

			result, repoErr := NewFaqReadRepository(mockedClient).Exists(context.TODO(), TestFaqID)

			tt.assertResult(t, tt.want, result, repoErr)
		})
	}
}
