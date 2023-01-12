package useCases

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/ssoql/faq-service/internal/useCases/repositories"
	"github.com/ssoql/faq-service/internal/useCases/repositories/mocks"
	"github.com/ssoql/faq-service/utils/apiErrors"
)

type deleteFaqTest struct{}

func (actualTest *deleteFaqTest) createRepositorySuccessMock(t *testing.T) repositories.FaqWriteRepository {
	r := mocks.NewFaqWriteRepository(t)
	r.On("Delete", mock.Anything, mock.Anything).Return(error(nil))

	return r
}

func (actualTest *deleteFaqTest) createRepositoryFailureMock(t *testing.T) repositories.FaqWriteRepository {
	r := mocks.NewFaqWriteRepository(t)
	r.On("Delete", mock.Anything, mock.Anything).Return(
		apiErrors.NewInternalServerError(
			"error when tying to fetch faq",
			errors.New("database error"),
		))

	return r
}

func Test_deleteFaqUseCase_Handle(t *testing.T) {
	type args struct {
		ctx   context.Context
		faqID int64
	}

	params := args{
		ctx:   context.Background(),
		faqID: 1,
	}
	actualTest := deleteFaqTest{}

	cases := []struct {
		name         string
		args         args
		repository   func(t *testing.T) repositories.FaqWriteRepository
		assertResult func(t *testing.T, err error)
	}{
		{
			name:       "success",
			repository: actualTest.createRepositorySuccessMock,
			args:       params,
			assertResult: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name:       "failure",
			repository: actualTest.createRepositoryFailureMock,
			args:       params,
			assertResult: func(t *testing.T, err error) {
				require.ErrorContains(t, err, "database error")
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			u := &deleteFaqUseCase{
				db: tt.repository(t),
			}
			err := u.Handle(tt.args.ctx, tt.args.faqID)

			tt.assertResult(t, err)
		})
	}
}
