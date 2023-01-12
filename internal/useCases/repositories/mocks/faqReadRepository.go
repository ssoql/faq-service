// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	apiErrors "github.com/ssoql/faq-service/utils/apiErrors"

	entities "github.com/ssoql/faq-service/internal/app/entities"

	mock "github.com/stretchr/testify/mock"
)

// FaqReadRepository is an autogenerated mock type for the FaqReadRepository type
type FaqReadRepository struct {
	mock.Mock
}

// Exists provides a mock function with given fields: ctx, id
func (_m *FaqReadRepository) Exists(ctx context.Context, id int64) (bool, apiErrors.ApiError) {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int64) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 apiErrors.ApiError
	if rf, ok := ret.Get(1).(func(context.Context, int64) apiErrors.ApiError); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apiErrors.ApiError)
		}
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx, page, pageSize
func (_m *FaqReadRepository) GetAll(ctx context.Context, page int, pageSize int) (*entities.Faqs, apiErrors.ApiError) {
	ret := _m.Called(ctx, page, pageSize)

	var r0 *entities.Faqs
	if rf, ok := ret.Get(0).(func(context.Context, int, int) *entities.Faqs); ok {
		r0 = rf(ctx, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Faqs)
		}
	}

	var r1 apiErrors.ApiError
	if rf, ok := ret.Get(1).(func(context.Context, int, int) apiErrors.ApiError); ok {
		r1 = rf(ctx, page, pageSize)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apiErrors.ApiError)
		}
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *FaqReadRepository) GetByID(ctx context.Context, id int64) (*entities.Faq, apiErrors.ApiError) {
	ret := _m.Called(ctx, id)

	var r0 *entities.Faq
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entities.Faq); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Faq)
		}
	}

	var r1 apiErrors.ApiError
	if rf, ok := ret.Get(1).(func(context.Context, int64) apiErrors.ApiError); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apiErrors.ApiError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewFaqReadRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewFaqReadRepository creates a new instance of FaqReadRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFaqReadRepository(t mockConstructorTestingTNewFaqReadRepository) *FaqReadRepository {
	mock := &FaqReadRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}