package useCases

import (
	"context"
	"errors"
	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/global"
	"github.com/ssoql/faq-service/internal/useCases/repositories"
	"github.com/ssoql/faq-service/internal/useCases/repositories/mocks"
	"github.com/ssoql/faq-service/utils/apiErrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

type updateFaqTest struct{}

func (actualTest *updateFaqTest) createWriteRepositorySuccessMock(t *testing.T) repositories.FaqWriteRepository {
	r := mocks.NewFaqWriteRepository(t)
	r.On("Update", mock.Anything, mock.Anything).Return(error(nil))

	return r
}

func (actualTest *updateFaqTest) createWriteRepositoryFailureMock(t *testing.T) repositories.FaqWriteRepository {
	r := mocks.NewFaqWriteRepository(t)
	r.On("Update", mock.Anything, mock.Anything).Return(
		apiErrors.NewInternalServerError(
			"error when tying to save faq",
			errors.New("database error"),
		),
	)

	return r
}

func (actualTest *updateFaqTest) createReadRepositorySuccessMock(t *testing.T) repositories.FaqReadRepository {
	r := mocks.NewFaqReadRepository(t)
	r.On("Exists", mock.Anything, mock.Anything).Return(true, error(nil))

	return r
}

func (actualTest *updateFaqTest) createReadRepositoryNotExistMock(t *testing.T) repositories.FaqReadRepository {
	r := mocks.NewFaqReadRepository(t)
	r.On("Exists", mock.Anything, mock.Anything).Return(false, error(nil))

	return r
}

func (actualTest *updateFaqTest) createReadRepositoryFailureMock(t *testing.T) repositories.FaqReadRepository {
	r := mocks.NewFaqReadRepository(t)
	r.On("Exists", mock.Anything, mock.Anything).Return(
		false,
		apiErrors.NewInternalServerError(
			"error when tying to save faq",
			errors.New("database error"),
		),
	)

	return r
}

func Test_updateFaqUseCase_Handle(t *testing.T) {
	type args struct {
		ctx        context.Context
		faqID      int64
		question   string
		answer     string
		isShutdown bool
	}

	params := args{
		ctx:        context.Background(),
		faqID:      int64(1),
		question:   "?",
		answer:     "!",
		isShutdown: false,
	}

	actualTest := updateFaqTest{}

	cases := []struct {
		name            string
		args            args
		readRepository  func(t *testing.T) repositories.FaqReadRepository
		writeRepository func(t *testing.T) repositories.FaqWriteRepository
		wantResult      *entities.Faq
		assertResult    func(t *testing.T, err error, getResult, wantResult *entities.Faq)
	}{
		{
			name:            "success",
			readRepository:  actualTest.createReadRepositorySuccessMock,
			writeRepository: actualTest.createWriteRepositorySuccessMock,
			args:            params,
			wantResult: &entities.Faq{
				Id:       params.faqID,
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
			name:            "update-failure",
			readRepository:  actualTest.createReadRepositorySuccessMock,
			writeRepository: actualTest.createWriteRepositoryFailureMock,
			args:            params,
			wantResult:      &entities.Faq{},
			assertResult: func(t *testing.T, err error, getResult, wantResult *entities.Faq) {
				require.ErrorContains(t, err, "database error")
			},
		},
		{
			name:           "check-if-exists-failure",
			readRepository: actualTest.createReadRepositoryFailureMock,
			writeRepository: func(t *testing.T) repositories.FaqWriteRepository {
				return mocks.NewFaqWriteRepository(t)
			},
			args:       params,
			wantResult: &entities.Faq{},
			assertResult: func(t *testing.T, err error, getResult, wantResult *entities.Faq) {
				require.ErrorContains(t, err, "database error")
			},
		},
		{
			name:           "not-exists",
			readRepository: actualTest.createReadRepositoryNotExistMock,
			writeRepository: func(t *testing.T) repositories.FaqWriteRepository {
				return mocks.NewFaqWriteRepository(t)
			},
			args:       params,
			wantResult: &entities.Faq{},
			assertResult: func(t *testing.T, err error, getResult, wantResult *entities.Faq) {
				require.ErrorContains(t, err, "faq with given id does not exist")
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			shutdownChan := make(chan os.Signal)

			defer close(shutdownChan)

			ctx := context.TODO()
			if tt.args.ctx != nil {
				ctx = tt.args.ctx
			}

			ctx = context.WithValue(ctx, global.ShutdownSignal, shutdownChan)
			// send signal to call graceful shutdown
			if tt.args.isShutdown {
				go func() {
					time.Sleep(20 * time.Millisecond)
					shutdownChan <- testShutdownSig{}
				}()
			}

			u := &updateFaqUseCase{
				dbWrite: tt.writeRepository(t),
				dbRead:  tt.readRepository(t),
			}

			result, err := u.Handle(ctx, tt.args.faqID, tt.args.question, tt.args.answer)

			tt.assertResult(t, err, result, tt.wantResult)
		})
	}
}
