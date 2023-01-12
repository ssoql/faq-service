package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/faq-service/internal/useCases"
	"github.com/ssoql/faq-service/utils/apiErrors"
	"net/http"
)

type FaqCreateInput struct {
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

type faqCreateHandler struct {
	useCase useCases.SaveFaqUseCase
}

func NewFaqCreateHandler(saveFaqUseCase useCases.SaveFaqUseCase) *faqCreateHandler {
	return &faqCreateHandler{useCase: saveFaqUseCase}
}

func (h *faqCreateHandler) Handle(c *gin.Context) {
	faqInput, err := retrieveValidInput(c)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	result, err := h.useCase.Handle(c, faqInput.Question, faqInput.Answer)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func retrieveValidInput(c *gin.Context) (*FaqCreateInput, apiErrors.ApiError) {
	var input FaqCreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, apiErrors.NewBadRequestError("invalid json data")
	}
	// todo add validation of parameters

	return &input, nil
}
