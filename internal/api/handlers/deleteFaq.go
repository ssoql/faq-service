package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/ssoql/faq-service/internal/useCases"
)

type faqDeleteHandler struct {
	useCase useCases.DeleteFaqUseCase
}

func NewFaqDeleteHandler(deleteFaqUseCase useCases.DeleteFaqUseCase) *faqDeleteHandler {
	return &faqDeleteHandler{useCase: deleteFaqUseCase}
}

func (h *faqDeleteHandler) Handle(c *gin.Context) {
	faqID, err := retrieveValidParameterID(c)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := h.useCase.Handle(c, faqID); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
