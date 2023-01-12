package infrastructure

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ssoql/faq-service/internal/infrastructure/db/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestFaqReadRepository(t *testing.T) {
	type args struct {
		faqID int64
	}
	tests := []struct {
		name        string
		args        args
		queriesMock func(t *testing.T, mockedDb sqlmock.Sqlmock)
		want        *faqWriteRepository
	}{
		{
			name: "success",
			args: args{
				faqID: int64(1),
			},
			queriesMock: func(t *testing.T, mockedDb sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "uniq_hash", "question", "answer"}).
					AddRow(
						int64(1),
						"35179a54ea587953021400eb0cd23201",
						"how are you?",
						"i am fine and you?")

				mockedDb.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM ` + "`faqs`" + ` WHERE id = ? ORDER BY ` + "`faqs`.`id`" + ` LIMIT 1`)).
					WithArgs(1).
					WillReturnRows(rows)
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
			tt.queriesMock(t, mockedDB)

			result, err := NewFaqReadRepository(mockedClient).GetByID(context.TODO(), tt.args.faqID)

			require.NoError(t, err)
			assert.Equal(t, tt.args.faqID, result.Id)
			assert.NotEmpty(t, result.UniqHash)
			assert.NotEmpty(t, result.Question)
			assert.NotEmpty(t, result.Answer)
		})
	}
}
