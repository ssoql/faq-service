package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ssoql/faq-service/internal/global"
	"github.com/ssoql/faq-service/utils/apiErrors"
)

type Pagination struct {
	pageNumber int
	pageSize   int
}

// InputValidationError implements AggregatedError and defines error for input errors.
type ValidationError struct {
	errors []error
}

func (e *ValidationError) Error() string {
	return Join(e.errors)
}

func (e *ValidationError) Slice() []error {
	return e.errors
}

func NewValidationError(errors ...error) *ValidationError {
	return &ValidationError{errors: errors}
}

func NewPagination() *Pagination {
	return &Pagination{pageNumber: 1, pageSize: 1}
}

func Join[T error](errs []T) string {
	builder := strings.Builder{}
	builder.WriteRune('[')

	msgs := make([]string, 0)
	for _, err := range errs {
		msgs = append(msgs, err.Error())
	}
	builder.WriteString(strings.Join(msgs, ", "))

	builder.WriteRune(']')

	return builder.String()
}

func retrieveValidParameterID(c *gin.Context) (int64, apiErrors.ApiError) {
	id, userErr := strconv.ParseInt(c.Param(global.ParameterFaqID), 10, 64)
	if userErr != nil {
		return 0, apiErrors.NewBadRequestError("id must be a number")
	}

	if id < 1 {
		return 0, apiErrors.NewBadRequestError("id must be greater than 0")
	}

	return id, nil
}

func validateGreaterZeroFunc(paramValue int, paramName string) func() error {
	return func() error {
		if paramValue < 1 {
			return fmt.Errorf("%s must be greater than 0", paramName)
		}
		return nil
	}
}

func validateIsEmptyOrNumericFunc(paramValue *string, paramName string) func() error {
	return func() error {
		if *paramValue == "" {
			return nil
		}
		_, userErr := strconv.Atoi(*paramValue)
		if userErr != nil {
			return fmt.Errorf("%s must be a number", paramName)
		}

		return nil
	}
}

// ExecuteValidation executes given validators without parameters. If no validation errors occurred it returns empty slice.
func ExecuteValidation(validators ...func() error) []error {
	allErrors := make([]error, 0)

	for _, validator := range validators {
		if err := validator(); err != nil {
			allErrors = append(allErrors, err)
		}
	}

	return allErrors
}

func retrieveValidPagination(c *gin.Context) (*Pagination, apiErrors.ApiError) {
	var errs []error

	rawPageNumber := c.Query(global.ParameterPage)
	rawPageSize := c.Query(global.ParameterPageSize)

	errs = ExecuteValidation(
		validateIsEmptyOrNumericFunc(&rawPageNumber, global.ParameterPage),
		validateIsEmptyOrNumericFunc(&rawPageSize, global.ParameterPageSize),
	)
	if len(errs) > 0 {
		return &Pagination{}, apiErrors.NewBadRequestError(NewValidationError(errs...).Error())
	}

	pagination := NewPagination()

	if rawPageNumber != "" {
		pagination.pageNumber, _ = strconv.Atoi(rawPageNumber)
	}

	if rawPageSize != "" {
		pagination.pageSize, _ = strconv.Atoi(rawPageSize)
	}

	errs = ExecuteValidation(
		validateGreaterZeroFunc(pagination.pageNumber, global.ParameterPage),
		validateGreaterZeroFunc(pagination.pageSize, global.ParameterPageSize),
	)
	if len(errs) > 0 {
		return &Pagination{}, apiErrors.NewBadRequestError(NewValidationError(errs...).Error())
	}

	return pagination, nil
}
