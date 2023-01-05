package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/faq-service/internal/useCases"
	"github.com/ssoql/faq-service/utils/api_errors"
	"net/http"
)

type FaqUpdateInput struct {
	*FaqCreateInput
	ID int64
}

type faqUpdateHandler struct {
	useCase useCases.UpdateFaqUseCase
}

func NewFaqUpdateHandler(updateFaqUseCase useCases.UpdateFaqUseCase) *faqUpdateHandler {
	return &faqUpdateHandler{useCase: updateFaqUseCase}
}

func (h *faqUpdateHandler) Handle(c *gin.Context) {
	faqInput, err := retrieveValidUpdateInput(c)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	//result, err := h.useCase.Handle(c, faqInput.Question, faqInput.Answer)
	//if err != nil {
	//	c.JSON(err.Status(), err)
	//	return
	//}

	c.JSON(http.StatusCreated, faqInput)
}

func retrieveValidUpdateInput(c *gin.Context) (*FaqUpdateInput, api_errors.ApiError) {
	var input FaqUpdateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, api_errors.NewBadRequestError("invalid json data")
	}
	// todo add validation of parameters
	id, err := retrieveValidParameterID(c)
	if err != nil {
		return nil, err
	}

	input.ID = id

	return &input, nil
}
