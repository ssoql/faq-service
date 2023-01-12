package useCases

import (
	"context"
	"errors"
	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/useCases/repositories"
	"github.com/ssoql/faq-service/internal/useCases/repositories/mocks"
	"github.com/ssoql/faq-service/utils/apiErrors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

type getFaqsTest struct{}

func (actualTest *getFaqsTest) createRepositorySuccessMock(t *testing.T) repositories.FaqReadRepository {
	r := mocks.NewFaqReadRepository(t)
	r.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(&entities.Faqs{{Id: 1}}, error(nil))

	return r
}

func (actualTest *getFaqsTest) createRepositoryFailureMock(t *testing.T) repositories.FaqReadRepository {
	r := mocks.NewFaqReadRepository(t)
	r.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(
		&entities.Faqs{},
		apiErrors.NewInternalServerError(
			"error when tying to fetch faq",
			errors.New("database error"),
		))

	return r
}

func Test_getFaqsUseCase_Handle(t *testing.T) {
	type args struct {
		ctx        context.Context
		pageNumber int
		pageSize   int
	}
	params := args{
		ctx:        context.Background(),
		pageNumber: 1,
		pageSize:   1,
	}

	actualTest := getFaqsTest{}

	cases := []struct {
		name         string
		args         args
		repository   func(t *testing.T) repositories.FaqReadRepository
		assertResult func(t *testing.T, err error, data *entities.Faqs)
	}{
		{
			name:       "success",
			repository: actualTest.createRepositorySuccessMock,
			args:       params,
			assertResult: func(t *testing.T, err error, data *entities.Faqs) {
				require.NoError(t, err)
				require.Greater(t, len(*data), 0)
			},
		},
		{
			name:       "failure",
			repository: actualTest.createRepositoryFailureMock,
			args:       params,
			assertResult: func(t *testing.T, err error, data *entities.Faqs) {
				require.ErrorContains(t, err, "database error")
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			u := &getFaqsUseCase{
				db: tt.repository(t),
			}
			result, err := u.Handle(tt.args.ctx, tt.args.pageNumber, tt.args.pageSize)
			tt.assertResult(t, err, result)
		})
	}
}
