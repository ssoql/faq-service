package useCases

import (
	"context"
	"errors"
	"github.com/ssoql/faq-service/internal/global"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

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

func (actualTest *saveFaqTest) createRepositoryShutdownMock(t *testing.T) repositories.FaqWriteRepository {
	r := mocks.NewFaqWriteRepository(t)
	r.On("Insert", mock.Anything, mock.Anything, mock.Anything).Return(
		func(ctx context.Context, faq *entities.Faq) apiErrors.ApiError {
			// wait for shutdown signal
			time.Sleep(200 * time.Millisecond)

			return apiErrors.NewInternalServerError(
				"error when tying to save faq",
				ctx.Err(),
			)
		},
	)

	return r
}

func Test_saveFaqUseCase_Handle(t *testing.T) {

	type args struct {
		ctx        context.Context
		question   string
		answer     string
		isShutdown bool
	}

	params := args{
		ctx:        context.Background(),
		question:   "?",
		answer:     "!",
		isShutdown: false,
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
			name:       "context-canceled",
			repository: actualTest.createRepositoryShutdownMock,
			args: func() args {
				params.isShutdown = true
				return params
			}(),
			wantResult: &entities.Faq{},
			assertResult: func(t *testing.T, err error, getResult, wantResult *entities.Faq) {
				require.ErrorContains(t, err, "context canceled")
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

			ctx := context.TODO()
			if tt.args.ctx != nil {
				ctx = tt.args.ctx
			}

			// send signal to call graceful shutdown
			if tt.args.isShutdown {
				shutdownChan := make(chan os.Signal)
				ctx = context.WithValue(ctx, global.ShutdownSignal, shutdownChan)

				go func() {
					defer close(shutdownChan)

					time.Sleep(20 * time.Millisecond)
					shutdownChan <- testShutdownSig{}
				}()
			}

			u := &saveFaqUseCase{
				db: tt.repository(t),
			}

			result, err := u.Handle(ctx, tt.args.question, tt.args.answer)

			tt.assertResult(t, err, result, tt.wantResult)
		})
	}
}
