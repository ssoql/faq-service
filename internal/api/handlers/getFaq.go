package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/faq-service/internal/useCases"
	"net/http"
)

type faqGetHandler struct {
	useCase useCases.GetFaqUseCase
}

type getFaqParameters struct {
	Id int64
}

func NewFaqGetHandler(getFaqUseCase useCases.GetFaqUseCase) *faqGetHandler {
	return &faqGetHandler{useCase: getFaqUseCase}
}

func (h *faqGetHandler) Handle(c *gin.Context) {
	faqID, err := retrieveValidParameterID(c)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	result, err := h.useCase.Handle(c, faqID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}
