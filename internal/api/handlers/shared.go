package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ssoql/faq-service/internal/global"
	"github.com/ssoql/faq-service/utils/api_errors"
)

func retrieveValidParameterID(c *gin.Context) (int64, api_errors.ApiError) {
	id, userErr := strconv.ParseInt(c.Param(global.ParameterFaqID), 10, 64)
	if userErr != nil {
		return 0, api_errors.NewBadRequestError("id must be a number")
	}

	if id < 1 {
		return 0, api_errors.NewBadRequestError("id must be greater than 0")
	}

	return id, nil
}
