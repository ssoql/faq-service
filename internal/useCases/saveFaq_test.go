package useCases

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/useCases/repositories"
	"github.com/ssoql/faq-service/internal/useCases/repositories/mocks"
	"github.com/ssoql/faq-service/utils/apiErrors"
)

type saveFaqTest struct{}

func (actualTest *saveFaqTest) createRepositorySuccessMock(t *testing.T) repositories.FaqWriteRepository {
	r := mocks.NewFaqWriteRepository(t)
	r.On("Insert", mock.Anything, mock.Anything).Return(func(ctx context.Context, faq *entities.Faq) apiErrors.ApiError {
		faq.Id = 1
		return nil
	})

	return r
}

func (actualTest *saveFaqTest) createRepositoryFailureMock(t *testing.T) repositories.FaqWriteRepository {
	r := mocks.NewFaqWriteRepository(t)
	r.On("Insert", mock.Anything, mock.Anything, mock.Anything).Return(
		apiErrors.NewInternalServerError(
			"error when tying to save faq",
			errors.New("database error"),
		),
	)

	return r
}

func Test_saveFaqUseCase_Handle(t *testing.T) {

	type args struct {
		ctx      context.Context
		question string
		answer   string
	}

	params := args{
		ctx:      context.Background(),
		question: "?",
		answer:   "!",
	}

	actualTest := saveFaqTest{}

	cases := []struct {
		name         string
		args         args
		repository   func(t *testing.T) repositories.FaqWriteRepository
		wantResult   *entities.Faq
		assertResult func(t *testing.T, err error, getResult, wantResult *entities.Faq)
	}{
		{
			name:       "success",
			repository: actualTest.createRepositorySuccessMock,
			args:       params,
			wantResult: &entities.Faq{
				Id:       int64(1),
				Question: params.question,
				Answer:   params.answer,
			},
			assertResult: func(t *testing.T, err error, getResult, wantResult *entities.Faq) {
				require.NoError(t, err)
				assert.Equal(t, wantResult.Id, getResult.Id)
				assert.Equal(t, wantResult.Question, getResult.Question)
				assert.Equal(t, wantResult.Answer, getResult.Answer)
			},
		},
		{
			name:       "failure",
			repository: actualTest.createRepositoryFailureMock,
			args:       params,
			wantResult: &entities.Faq{},
			assertResult: func(t *testing.T, err error, getResult, wantResult *entities.Faq) {
				require.ErrorContains(t, err, "database error")
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			u := &saveFaqUseCase{
				db: tt.repository(t),
			}
			result, err := u.Handle(tt.args.ctx, tt.args.question, tt.args.answer)

			tt.assertResult(t, err, result, tt.wantResult)
		})
	}
}
