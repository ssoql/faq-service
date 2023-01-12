package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/faq-service/internal/useCases"
	"net/http"
)

type faqsGetHandler struct {
	useCase useCases.GetFaqsUseCase
}

func NewFaqsGetHandler(getFaqsUseCase useCases.GetFaqsUseCase) *faqsGetHandler {
	return &faqsGetHandler{useCase: getFaqsUseCase}
}

func (h *faqsGetHandler) Handle(c *gin.Context) {
	pagination, err := retrieveValidPagination(c)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	result, err := h.useCase.Handle(c, pagination.pageNumber, pagination.pageSize)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}
