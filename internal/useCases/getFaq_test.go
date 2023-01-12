package useCases

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/useCases/repositories"
	"github.com/ssoql/faq-service/internal/useCases/repositories/mocks"
	"github.com/ssoql/faq-service/utils/apiErrors"
	"github.com/stretchr/testify/mock"
)

type getFaqTest struct{}

func (actualTest *getFaqTest) createRepositorySuccessMock(t *testing.T) repositories.FaqReadRepository {
	r := mocks.NewFaqReadRepository(t)
	r.On("GetByID", mock.Anything, mock.Anything).Return(&entities.Faq{Id: 1}, error(nil))

	return r
}

func (actualTest *getFaqTest) createRepositoryFailureMock(t *testing.T) repositories.FaqReadRepository {
	r := mocks.NewFaqReadRepository(t)
	r.On("GetByID", mock.Anything, mock.Anything).Return(
		&entities.Faq{},
		apiErrors.NewInternalServerError(
			"error when tying to fetch faq",
			errors.New("database error"),
		))

	return r
}

func Test_getFaqUseCase_Handle(t *testing.T) {

	type args struct {
		ctx   context.Context
		faqID int64
	}

	params := args{
		ctx:   context.Background(),
		faqID: 1,
	}

	actualTest := getFaqTest{}

	cases := []struct {
		name         string
		args         args
		repository   func(t *testing.T) repositories.FaqReadRepository
		assertResult func(t *testing.T, err error, data *entities.Faq)
	}{
		{
			name:       "success",
			repository: actualTest.createRepositorySuccessMock,
			args:       params,
			assertResult: func(t *testing.T, err error, data *entities.Faq) {
				require.NoError(t, err)
				require.Equal(t, int64(1), data.Id)
			},
		},
		{
			name:       "failure",
			repository: actualTest.createRepositoryFailureMock,
			args:       params,
			assertResult: func(t *testing.T, err error, data *entities.Faq) {
				require.ErrorContains(t, err, "database error")
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			u := &getFaqUseCase{
				db: tt.repository(t),
			}
			result, err := u.Handle(tt.args.ctx, tt.args.faqID)

			tt.assertResult(t, err, result)
		})
	}
}
